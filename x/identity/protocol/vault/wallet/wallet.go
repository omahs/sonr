package wallet

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/sonrhq/core/x/identity/protocol/vault/account"
	v1 "github.com/sonrhq/core/x/identity/types/vault/v1"
)

// `Wallet` is an interface that has a method `WalletConfig` that returns a `*v1.WalletConfig` and a
// method `CreateAccount` that takes a `string`, a `string`, and a `string` and returns an `error`.
// @property WalletConfig - This is the configuration of the wallet.
// @property {error} CreateAccount - Creates a new account
// @property GetAccount - Returns the account with the given name.
// @property PrimaryAccount - The primary account is the account that is used to sign transactions.
// @property ListAccounts - Returns a list of all accounts in the wallet
type Wallet interface {
	// The wallet configuration
	WalletConfig() *v1.WalletConfig

	// Creates a new account
	CreateAccount(name string, addrPrefix string, networkName string) error

	// Gets an account by name
	GetAccount(name string) (account.WalletAccount, error)

	// Gets Primary account
	PrimaryAccount() (account.WalletAccount, error)

	// Gets all accounts
	ListAccounts() ([]account.WalletAccount, error)

	// Signs a transaction for Cosmos compatible blockchains
	SendTx(memo string, msgs ...sdk.Msg) (*sdk.TxResponse, error)
}

// `walletImpl` is a struct that has a single field, `walletConfig`, which is a pointer to a
// `v1.WalletConfig` struct.
// @property walletConfig - The wallet configuration
type walletImpl struct {
	// The wallet configuration
	walletConfig *v1.WalletConfig

	// The TxBuilder
	cctx client.Context
}

// `NewWallet` creates a new wallet with a default root account
func NewWallet(cctx client.Context) (Wallet, error) {
	// The default shards that are added to the MPC wallet
	rootAcc, err := account.NewAccount("Primary", "snr", "Sonr")
	if err != nil {
		return nil, err
	}
	conf := v1.NewWalletConfigFromRootAccount(rootAcc.AccountConfig())
	return &walletImpl{
		walletConfig: conf,
		cctx:         cctx,
	}, nil
}

// `NewWalletFromConfig` takes a `WalletConfig` and returns a `Wallet` and an error
func NewWalletFromConfig(walletConf *v1.WalletConfig) (Wallet, error) {
	return &walletImpl{
		walletConfig: walletConf,
	}, nil
}

// Returning the wallet configuration.
func (w *walletImpl) WalletConfig() *v1.WalletConfig {
	return w.walletConfig
}

// Creating a new account.
func (w *walletImpl) CreateAccount(name string, addrPrefix string, networkName string) error {
	// The default shards that are added to the MPC wallet
	rootAcc, err := w.PrimaryAccount()
	if err != nil {
		return err
	}
	acc, err := rootAcc.Bip32Derive(name, uint32(len(w.walletConfig.Accounts)), addrPrefix, networkName)
	if err != nil {
		return err
	}
	w.walletConfig.Accounts[name] = acc.AccountConfig()
	return nil
}

// Returning the account.WalletAccount and error.
func (w *walletImpl) GetAccount(name string) (account.WalletAccount, error) {
	accConf, ok := w.walletConfig.Accounts[strings.ToLower(name)]
	if !ok {
		return nil, errors.New("Account not found")
	}
	return account.NewAccountFromConfig(accConf)
}

// Returning a list of accounts.
func (w *walletImpl) ListAccounts() ([]account.WalletAccount, error) {
	accs := make([]account.WalletAccount, 0, len(w.walletConfig.Accounts))
	for _, accConf := range w.walletConfig.Accounts {
		acc, err := account.NewAccountFromConfig(accConf)
		if err != nil {
			return nil, err
		}
		accs = append(accs, acc)
	}
	return accs, nil
}

// Returning the primary account.
func (w *walletImpl) PrimaryAccount() (account.WalletAccount, error) {
	return w.GetAccount("Primary")
}

// Signing a transaction.
func (w *walletImpl) SendTx(memo string, msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	req, err := w.buildBroadcastTx(memo, msgs...)
	if err != nil {
		return nil, err
	}
	cl, err := client.NewClientFromNode(w.cctx.NodeURI)
	if err != nil {
		return nil, err
	}
	res, err := cl.BroadcastTxCommit(context.Background(), req)
	if err != nil {
		return nil, err
	}
	txres := &sdk.TxResponse{
		Height:    res.Height,
		TxHash:    string(res.DeliverTx.Data),
		Codespace: res.DeliverTx.Codespace,
		Code:      res.DeliverTx.Code,
		RawLog:    res.DeliverTx.Log,
	}
	if txres.Code != 0 {
		return nil, fmt.Errorf("tx failed: %s", txres.RawLog)
	}
	return txres, nil
}

// Building a transaction from the given inputs.
func (w *walletImpl) buildBroadcastTx(memo string, msgs ...sdk.Msg) ([]byte, error) {
	prim, err := w.PrimaryAccount()
	if err != nil {
		return nil, err
	}
	auxData, err := prim.SignTxAux(msgs...)
	if err != nil {
		return nil, err
	}
	txBuilder := w.cctx.TxConfig.NewTxBuilder()
	txBuilder.SetMemo(memo)
	txBuilder.SetMsgs(msgs...)
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(100))))
	txBuilder.AddAuxSignerData(auxData)
	txBuilder.SetGasLimit(200000)
	return w.cctx.TxConfig.TxEncoder()(txBuilder.GetTx())
}

//
// Helper functions
//

// makeTxBody builds a transaction from the given inputs.
func makeTxBody(note string, msgs ...sdk.Msg) (*txtypes.TxBody, error) {
	// func BuildTx(w *crypto.MPCWallet, msgs ...sdk.Msg) (*txtypes.TxBody, error) {
	// Create Any for each message
	anyMsgs := make([]*codectypes.Any, len(msgs))
	for i, m := range msgs {
		msg, err := codectypes.NewAnyWithValue(m)
		if err != nil {
			return nil, err
		}
		anyMsgs[i] = msg
	}

	// Create TXRaw and Marshal
	txBody := txtypes.TxBody{
		Messages: anyMsgs,
		Memo:     note,
	}
	return &txBody, nil
}

// createRawTxBytes is a helper function to create a raw raw transaction and Marshal it to bytes
func createRawTxBytes(txBody *txtypes.TxBody, sig []byte, authInfo *txtypes.AuthInfo) ([]byte, error) {
	// Serialize the tx body
	txBytes, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}

	// Serialize the authInfo
	authInfoBytes, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}

	// Create a signature list and append the signature
	sigList := make([][]byte, 1)
	sigList[0] = sig

	// Create Raw TX
	txRaw := &txtypes.TxRaw{
		BodyBytes:     txBytes,
		AuthInfoBytes: authInfoBytes,
		Signatures:    sigList,
	}

	// Marshal the txRaw
	return txRaw.Marshal()
}

// getAuthInfoSingle returns the authentication information for the given message.
func getAuthInfoSingle(w account.WalletAccount, gas int) (*txtypes.AuthInfo, error) {
	// Build signerInfo parameters
	anyPubKey, err := codectypes.NewAnyWithValue(w.PubKey())
	if err != nil {
		return nil, err
	}

	// Create AuthInfo
	authInfo := txtypes.AuthInfo{
		SignerInfos: []*txtypes.SignerInfo{
			{
				PublicKey: anyPubKey,
				ModeInfo: &txtypes.ModeInfo{
					Sum: &txtypes.ModeInfo_Single_{
						Single: &txtypes.ModeInfo_Single{
							Mode: 3,
						},
					},
				},
			},
		},
		Fee: &txtypes.Fee{
			Amount:   sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(int64(gas)))),
			GasLimit: uint64(300000),
			Granter:  w.AccountConfig().Address,
		},
	}
	return &authInfo, nil
}

// It takes a transaction body and auth info, serializes them, and then creates a SignDoc object that
// contains the serialized transaction body and auth info, and the chain ID
func getSignDocBytes(authInfo *txtypes.AuthInfo, txBody *txtypes.TxBody) ([]byte, error) {
	// Serialize the transaction body.
	txBodyBz, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}

	// Serialize the auth info.
	authInfoBz, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}

	// Create SignDoc
	signDoc := &txtypes.SignDoc{
		BodyBytes:     txBodyBz,
		AuthInfoBytes: authInfoBz,
		ChainId:       "sonr",
	}
	return signDoc.Marshal()
}
