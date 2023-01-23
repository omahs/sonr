package types

import (
	"bytes"
	"errors"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/shengdoushi/base58"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	tmcrypto "github.com/tendermint/tendermint/crypto"
)

type (
	Address = tmcrypto.Address
)

//
// Constructors
//

// It takes a string, decodes it from base58, unmarshals it into a PubKey, and returns the PubKey
func NewPubKeyFromBase58(key string) (*PubKey, error) {
	bz, err := base58.Decode(key, base58.BitcoinAlphabet)
	if err != nil {
		return nil, err
	}
	pk := &PubKey{
		Key: bz,
	}
	return pk, nil
}

// NewPubKeyFromBytes takes a byte array and returns a PubKey
func NewPubKeyFromBytes(bz []byte) *PubKey {
	pk := &PubKey{}
	pk.Key = bz
	return pk
}

//
// CryptoTypes Implementation of PubKey interface
//

// Creating a new method called Address() that returns an Address type.
func (pk *PubKey) Address() Address {
	return tmcrypto.AddressHash(pk.Key)
}

// Returning the key in bytes.
func (pk *PubKey) Bytes() []byte {
	return pk.Key
}

// Verifying the signature of the message.
func (pk *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	pp := &curve.Secp256k1Point{}
	if err := pp.UnmarshalBinary(pk.Key); err != nil {
		return false
	}
	signature, err := deserializeSignature(sig)
	if err != nil {
		return false
	}
	return signature.Verify(pp, msg)
}

// Comparing the two keys.
func (pk *PubKey) Equals(other cryptotypes.PubKey) bool {
	if other == nil {
		return false
	}
	return bytes.Equal(pk.Key, other.Bytes())
}

// Returning the type of the key.
func (pk *PubKey) Type() string {
	return "secp256k1"
}

// SerializeSignature marshals an ECDSA signature to DER format for use with the CMP protocol
func serializeSignature(sig *ecdsa.Signature) ([]byte, error) {
	rBytes, err := sig.R.MarshalBinary()
	if err != nil {
		return nil, err
	}
	sBytes, err := sig.S.MarshalBinary()
	if err != nil {
		return nil, err
	}

	sigBytes := make([]byte, 65)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[65-len(sBytes):65], sBytes)
	return sigBytes, nil
}

// - The R and S values must be in the valid range for secp256k1 scalars:
//   - Negative values are rejected
//   - Zero is rejected
//   - Values greater than or equal to the secp256k1 group order are rejected
func deserializeSignature(sigStr []byte) (*ecdsa.Signature, error) {
	rBytes := sigStr[:33]
	sBytes := sigStr[33:65]

	sig := ecdsa.EmptySignature(curve.Secp256k1{})
	if err := sig.R.UnmarshalBinary(rBytes); err != nil {
		return nil, errors.New("malformed signature: R is not in the range [1, N-1]")
	}

	// S must be in the range [1, N-1].  Notice the check for the maximum number
	// of bytes is required because SetByteSlice truncates as noted in its
	// comment so it could otherwise fail to detect the overflow.
	if err := sig.S.UnmarshalBinary(sBytes); err != nil {
		return nil, errors.New("malformed signature: S is not in the range [1, N-1]")
	}

	// Create and return the signature.
	return &sig, nil
}
