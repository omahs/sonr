package protocol

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

const Bip44Prefix = "m"

type Purpose uint32

const (
	PurposeBIP44 Purpose = 44
)

type KeyShareParseResult struct {
	Purpose      Purpose
	CoinType     crypto.CoinType
	AccountName  string
	KeyShareName string
}

func ParseKeyShareName(name string) (*KeyShareParseResult, error) {
	// Keyshare name format is /{purpose}/{coin_type}/{account_name}/{keyshare_name}
	ptrs := strings.Split(name, "/")
	if len(ptrs) != 5 {
		return nil, fmt.Errorf("invalid keyshare name: %s, LENGTH: %v", name, len(ptrs))
	}

	purposeStr := ptrs[1]
	coinTypeStr := ptrs[2]
	accountName := ptrs[3]
	keyShareName := ptrs[4]

	var purpose Purpose
	switch purposeStr {
	case "44":
		purpose = PurposeBIP44
	default:
		return nil, fmt.Errorf("invalid purpose: %s", purposeStr)
	}

	coinType, err := strconv.Atoi(coinTypeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid coin type: %s", coinTypeStr)
	}
	ct := crypto.CoinTypeFromBipPath(int32(coinType))

	return &KeyShareParseResult{
		Purpose:      purpose,
		CoinType:     ct,
		AccountName:  accountName,
		KeyShareName: keyShareName,
	}, nil
}

// KeyShare is a type that interacts with a cmp.Config file located on disk.
type KeyShare interface {
	// Bip44 returns the bip44 path for the keyshare
	Bip44() string

	// Config returns the cmp.Config.
	Config() *cmp.Config

	// CoinType returns the coin type based on the account directories parent
	CoinType() crypto.CoinType

	// AccountName returns the account name based on the account directory name
	AccountName() string

	// KeyID returns the key id based on the keyshare file name
	KeyID() string

	// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
	Encrypt(credential *crypto.WebauthnCredential) error

	// Encrypt checks if the file at current path is encrypted and if not, encrypts it.
	Decrypt(credential *crypto.WebauthnCredential) error

	IsEncrypted() bool
}

// keyShare is a type that interacts with a cmp.Config file located on disk.
type keyShare struct {
	bytes    []byte
	name     string
	lastUsed uint32
}

// Keyshare name format is /{purpose}/{coin_type}/{account_name}/{keyshare_name}
func NewKeyshare(id string, bytes []byte, coinType crypto.CoinType, accName string) (KeyShare, error) {
	name := fmt.Sprintf("/%d/%d/%s/%s", PurposeBIP44, coinType.BipPath(), accName, id)
	conf := cmp.EmptyConfig(curve.Secp256k1{})
	err := conf.UnmarshalBinary(bytes)
	if err != nil {
		return nil, err
	}
	return &keyShare{
		bytes:    bytes,
		name:     name,
		lastUsed: uint32(time.Now().Unix()),
	}, nil
}

// Bip44 returns the bip44 path for the keyshare
func (ks *keyShare) Bip44() string {
	return filepath.Join(Bip44Prefix, ks.name)
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
	res, err := ParseKeyShareName(ks.name)
	if err != nil {
		panic(err)
	}
	return res.CoinType
}

// AccountName returns the account name based on the account directory name
func (ks *keyShare) AccountName() string {
		// Keyshare name format is /{purpose}/{coin_type}/{account_name}/{keyshare_name}
	ptrs := strings.Split(ks.name, "/")
	if len(ptrs) != 5 {
		panic("invalid keyshare name")
	}
	accountName := ptrs[3]
	return accountName
}

// KeyID returns the key id based on the keyshare file name
func (ks *keyShare) KeyID() string {
	// Keyshare name format is /{purpose}/{coin_type}/{account_name}/{keyshare_name}
	ptrs := strings.Split(ks.name, "/")
	if len(ptrs) != 5 {
		panic("invalid keyshare name")
	}
	keyShareName := ptrs[4]
	return keyShareName
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

// A Keyshare is encrypted if its name contains an apostrophe at the end.
func (ks *keyShare) IsEncrypted() bool {
	if ks.name == "vault" {
		return false
	}
	return strings.HasSuffix(ks.name, "'")
}

type AccountI map[string][]byte

type WalletI map[string]AccountI

func NewAccountI() AccountI {
	return make(AccountI, 0)
}

func (a AccountI) Add(key string, value []byte) {
	a[key] = value
}

func (a AccountI) Get(key string) []byte {
	return a[key]
}

func (a AccountI) Delete(key string) {
	delete(a, key)
}

func (a AccountI) Keys() []string {
	keys := make([]string, 0)
	for k := range a {
		keys = append(keys, k)
	}
	return keys
}

func (a AccountI) Values() [][]byte {
	values := make([][]byte, 0)
	for _, v := range a {
		values = append(values, v)
	}
	return values
}
