package tx

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/sonrhq/core/x/identity/protocol/vault/account"
)

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
