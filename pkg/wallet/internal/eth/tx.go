package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/sonrhq/core/pkg/wallet"
)

// EthereumTransaction represents the Ethereum transaction data
type EthereumTransaction struct {
	Nonce    uint64
	To       string
	Value    *big.Int
	GasLimit uint64
	GasPrice *big.Int
	Data     []byte
}

// SignEthereumTransaction signs an Ethereum transaction using the Sonr wallet account abstraction
func SignEthereumTransaction(wa wallet.Account, etx *EthereumTransaction) ([]byte, error) {
	// Serialize the Ethereum transaction data
	txData := types.NewTransaction(etx.Nonce, common.HexToAddress(etx.To), etx.Value, etx.GasLimit, etx.GasPrice, etx.Data)
	encodedTx, err := rlp.EncodeToBytes(txData)
	if err != nil {
		return nil, err
	}

	// Sign the serialized transaction data using the existing Sign method
	signedTx, err := wa.Sign(encodedTx)
	if err != nil {
		return nil, err
	}

	// Add the Ethereum-specific ECDSA recovery id (v) to the signature
	recId := byte(27)
	if crypto.VerifySignature(wa.PubKey().Bytes(), encodedTx, signedTx[:64]) {
		recId = byte(28)
	}

	signedTx[64] = recId - byte(27)

	return signedTx, nil
}
