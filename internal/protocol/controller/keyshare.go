package controller

import (
	"fmt"
	"strings"
	"time"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)


var ControllerTypeSystem = schema.MustTypeSystem(
	schema.SpawnStruct("keyShare",
		[]schema.StructField{
			schema.SpawnStructField("bytes", "Bytes", false, false),
			schema.SpawnStructField("name", "String", false, false),
			schema.SpawnStructField("lastUsed", "Int", false, false),
		},
		schema.SpawnStructRepresentationMap(nil),
	),
	schema.SpawnBytes("Bytes"),
	schema.SpawnString("String"),
	schema.SpawnInt("Int"),
)

// KeyShare is a type that interacts with a cmp.Config file located on disk.
type KeyShare interface {
	// Bytes returns the bytes of the keyshare file - the marshalled cmp.Config
	Bytes() []byte

	// Config returns the cmp.Config.
	Config() *cmp.Config

	// CoinType returns the coin type based on the account directories parent
	CoinType() crypto.CoinType

	// AccountName returns the account name based on the account directory name
	AccountName() string

	// Did returns the cid of the keyshare
	Did() string

	// PartyID returns the party id based on the keyshare file name
	PartyID() crypto.PartyID

	// PubKey returns the public key of the keyshare
	PubKey() *crypto.PubKey

	// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
	Encrypt(credential *crypto.WebauthnCredential) error

	// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
	Decrypt(credential *crypto.WebauthnCredential) error

	// IsEncrypted checks if the file at current path is encrypted.
	IsEncrypted() bool
}

// keyShare is a type that interacts with a cmp.Config file located on disk.
type keyShare struct {
	bytes    []byte
	name     string
	lastUsed uint32
}


type Foobar struct {
	Foo string
	Bar string
}

// Keyshare name format is a DID did:{coin_type}:{account_address}#ks-{keyshare_name}
func NewKeyshare(id string, bytes []byte, coinType crypto.CoinType, accName string) (KeyShare, error) {
	conf := cmp.EmptyConfig(curve.Secp256k1{})
	err := conf.UnmarshalBinary(bytes)
	if err != nil {
		return nil, err
	}

	ks := &keyShare{
		bytes:    bytes,
		lastUsed: uint32(time.Now().Unix()),
	}
	addr := coinType.FormatAddress(ks.PubKey())
	ks.name = fmt.Sprintf("did:%s:%s#ks-%s", coinType.DidMethod(), addr, string(conf.ID))
	return ks, nil
}

// LoadKeyshareFromStore loads a keyshare from a store.
func LoadKeyshareFromStore(key string, value []byte) (KeyShare, error) {
	ksr, err := ParseKeyShareDid(key)
	if err != nil {
		return nil, err
	}
	conf := cmp.EmptyConfig(curve.Secp256k1{})
	err = conf.UnmarshalBinary(value)
	if err != nil {
		return nil, err
	}

	return &keyShare{
		bytes:    value,
		name:     ksr.KeyShareName,
		lastUsed: uint32(time.Now().Unix()),
	}, nil
}

// Cid returns the cid of the keyshare
func (ks *keyShare) Did() string {
	return ks.name
}

// Bytes returns the bytes of the keyshare file - the marshalled cmp.Config
func (ks *keyShare) Bytes() []byte {
	return ks.bytes
}

// Config returns the cmp.Config.
func (ks *keyShare) Config() *cmp.Config {
	cnfg := cmp.EmptyConfig(curve.Secp256k1{})
	err := cnfg.UnmarshalBinary(ks.bytes)
	if err != nil {
		panic(err)
	}
	ks.lastUsed = uint32(time.Now().Unix())
	return cnfg
}

// CoinType returns the coin type based on the account directories parent
func (ks *keyShare) CoinType() crypto.CoinType {
	res, err := ParseKeyShareDid(ks.name)
	if err != nil {
		panic(err)
	}
	return res.CoinType
}

// AccountName returns the account name based on the account directory name
func (ks *keyShare) AccountName() string {
	res, err := ParseKeyShareDid(ks.name)
	if err != nil {
		panic(err)
	}
	return res.AccountName
}

// Keyshare name format is /{purpose}/{coin_type}/{account_name}/{keyshare_name}
func (ks *keyShare) KeyID() string {
	return ks.name
}

// PartyID returns the party id based on the keyshare file name
func (ks *keyShare) PartyID() crypto.PartyID {
	res, err := ParseKeyShareDid(ks.name)
	if err != nil {
		panic(err)
	}
	return crypto.PartyID(res.KeyShareName)
}

// PublicKey returns the public key of the keyshare
func (ks *keyShare) PubKey() *crypto.PubKey {
	skPP, ok := ks.Config().PublicPoint().(*curve.Secp256k1Point)
	if !ok {
		return nil
	}
	bz, err := skPP.MarshalBinary()
	if err != nil {
		return nil
	}
	return crypto.NewSecp256k1PubKey(bz)
}

// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
func (ks *keyShare) Encrypt(credential *crypto.WebauthnCredential) error {
	if ks.name == "vault" {
		return nil
	}
	enc, err := credential.Encrypt(ks.bytes)
	if err != nil {
		return err
	}
	ks.lastUsed = uint32(time.Now().Unix())
	ks.bytes = enc
	ks.name += "'" // encrypted keyshares have an apostrophe at the end
	return nil
}

// Decrypt checks if the file at current path is encrypted and if not, encrypts it.
func (ks *keyShare) Decrypt(credential *crypto.WebauthnCredential) error {
	if !ks.IsEncrypted() {
		return nil
	}

	dec, err := credential.Decrypt(ks.bytes)
	if err != nil {
		return err
	}
	ks.lastUsed = uint32(time.Now().Unix())
	ks.bytes = dec
	ks.name = strings.TrimSuffix(ks.name, "'") // remove the apostrophe
	return nil
}

// IPLD returns the ipld node of the keyshare
func (ks *keyShare) Block() blocks.Block {
	bz, err := ipld.Marshal(dagjson.Encode, ks, ControllerTypeSystem.TypeByName("keyShare"))
	if err != nil {
		panic(err)
	}
	return blocks.NewBlock(bz)
}

// A Keyshare is encrypted if its name contains an apostrophe at the end.
func (ks *keyShare) IsEncrypted() bool {
	if ks.name == "vault" {
		return false
	}
	return strings.HasSuffix(ks.name, "'")
}
