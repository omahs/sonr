package crypto

import (
	fmt "fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
)

// AllCoinTypes returns all the coin types.
func AllCoinTypes() []CoinType {
	return []CoinType{
		CoinType_CoinType_BITCOIN,
		CoinType_CoinType_ETHEREUM,
		CoinType_CoinType_LITECOIN,
		CoinType_CoinType_DOGE,
		CoinType_CoinType_SONR,
		CoinType_CoinType_COSMOS,
		CoinType_CoinType_FILECOIN,
		CoinType_CoinType_HNS,
		CoinType_CoinType_TESTNET,
		CoinType_CoinType_SOLANA,
		CoinType_CoinType_XRP,
	}
}

// CoinTypeFromAddrPrefix returns the CoinType from the string.
func CoinTypeFromAddrPrefix(str string) CoinType {
	coins := AllCoinTypes()
	for _, coin := range coins {
		if coin.AddrPrefix() == str {
			return coin
		}
	}
	return CoinType_CoinType_TESTNET
}

// CoinTypeFromBipPath returns the CoinType from the index.
func CoinTypeFromBipPath(i int32) CoinType {
	coins := AllCoinTypes()
	for _, coin := range coins {
		if coin.BipPath() == i {
			return coin
		}
	}
	return CoinType_CoinType_TESTNET
}

// CoinTypeFromDIDMethod returns the CoinType from the DID method.
func CoinTypeFromDidMethod(str string) CoinType {
	coins := AllCoinTypes()
	for _, coin := range coins {
		if strings.ToLower(coin.DidMethod()) == strings.ToLower(str) {
			return coin
		}
	}
	return CoinType_CoinType_TESTNET
}

// CoinTypeFromName returns the CoinType from the name.
func CoinTypeFromName(str string) CoinType {
	coins := AllCoinTypes()
	for _, coin := range coins {
		if strings.ToLower(coin.Name()) == strings.ToLower(str) {
			return coin
		}
	}
	return CoinType_CoinType_TESTNET
}

// CoinTypeFromTicker returns the CoinType from the symbol.
func CoinTypeFromTicker(str string) CoinType {
	coins := AllCoinTypes()
	for _, coin := range coins {
		if strings.ToLower(coin.Ticker()) == strings.ToLower(str) {
			return coin
		}
	}
	return CoinType_CoinType_TESTNET
}

// AddrPrefix returns the address prefix for the given coin type.
func (ct CoinType) AddrPrefix() string {
	switch ct {
	case CoinType_CoinType_BITCOIN:
		return "bc"
	case CoinType_CoinType_ETHEREUM:
		return "0x"
	case CoinType_CoinType_LITECOIN:
		return "ltc"
	case CoinType_CoinType_DOGE:
		return "doge"
	case CoinType_CoinType_SONR:
		return "idx"
	case CoinType_CoinType_COSMOS:
		return "cosmos"
	case CoinType_CoinType_FILECOIN:
		return "f"
	case CoinType_CoinType_HNS:
		return "hs"
	case CoinType_CoinType_TESTNET:
		return "test"
	case CoinType_CoinType_SOLANA:
		return "sol"
	case CoinType_CoinType_XRP:
		return "xrp"
	default:
		return "test"
	}
}

// BipPath returns the index for the given coin type.
func (ct CoinType) BipPath() int32 {
	switch ct {
	case CoinType_CoinType_BITCOIN:
		return 0
	case CoinType_CoinType_ETHEREUM:
		return 60
	case CoinType_CoinType_LITECOIN:
		return 2
	case CoinType_CoinType_DOGE:
		return 3
	case CoinType_CoinType_SONR:
		return 703
	case CoinType_CoinType_COSMOS:
		return 118
	case CoinType_CoinType_FILECOIN:
		return 461
	case CoinType_CoinType_HNS:
		return 5353
	case CoinType_CoinType_TESTNET:
		return 1
	case CoinType_CoinType_SOLANA:
		return 501
	case CoinType_CoinType_XRP:
		return 144
	default:
		return 1
	}
}

// Name returns the name for the given coin type.
func (ct CoinType) Name() string {
	switch ct {
	case CoinType_CoinType_BITCOIN:
		return "Bitcoin"
	case CoinType_CoinType_ETHEREUM:
		return "Ethereum"
	case CoinType_CoinType_LITECOIN:
		return "Litecoin"
	case CoinType_CoinType_DOGE:
		return "Dogecoin"
	case CoinType_CoinType_SONR:
		return "Sonr"
	case CoinType_CoinType_COSMOS:
		return "Cosmos"
	case CoinType_CoinType_FILECOIN:
		return "Filecoin"
	case CoinType_CoinType_HNS:
		return "Handshake"
	case CoinType_CoinType_SOLANA:
		return "Solana"
	case CoinType_CoinType_XRP:
		return "Ripple"
	default:
		return "Testnet"
	}
}

// Ticker returns the symbol for the given coin type for exchanges.
func (ct CoinType) Ticker() string {
	switch ct {
	case CoinType_CoinType_BITCOIN:
		return "BTC"
	case CoinType_CoinType_ETHEREUM:
		return "ETH"
	case CoinType_CoinType_LITECOIN:
		return "LTC"
	case CoinType_CoinType_DOGE:
		return "DOGE"
	case CoinType_CoinType_SONR:
		return "SNR"
	case CoinType_CoinType_COSMOS:
		return "ATOM"
	case CoinType_CoinType_FILECOIN:
		return "FIL"
	case CoinType_CoinType_HNS:
		return "HNS"
	case CoinType_CoinType_SOLANA:
		return "SOL"
	case CoinType_CoinType_XRP:
		return "XRP"
	default:
		return "TESTNET"
	}
}

// IsBitcoin returns true if the coin type is bitcoin.
func (c CoinType) IsBitcoin() bool {
	return c == CoinType_CoinType_BITCOIN
}

// IsCosmos returns true if the coin type is cosmos.
func (c CoinType) IsCosmos() bool {
	return c == CoinType_CoinType_COSMOS
}

// IsEthereum returns true if the coin type is ethereum.
func (c CoinType) IsEthereum() bool {
	return c == CoinType_CoinType_ETHEREUM
}

// IsFilecoin returns true if the coin type is filecoin.
func (c CoinType) IsFilecoin() bool {
	return c == CoinType_CoinType_FILECOIN
}

// IsHandshake returns true if the coin type is handshake.
func (c CoinType) IsHandshake() bool {
	return c == CoinType_CoinType_HNS
}

// IsLitecoin returns true if the coin type is litecoin.
func (c CoinType) IsLitecoin() bool {
	return c == CoinType_CoinType_LITECOIN
}

// IsSolana returns true if the coin type is solana.
func (c CoinType) IsSolana() bool {
	return c == CoinType_CoinType_SOLANA
}

// IsRipple returns true if the coin type is ripple.
func (c CoinType) IsRipple() bool {
	return c == CoinType_CoinType_XRP
}

// IsTestnet returns true if the coin type is testnet.
func (c CoinType) IsTestnet() bool {
	return c == CoinType_CoinType_TESTNET
}

// IsDogecoin returns true if the coin type is dogecoin.
func (c CoinType) IsDogecoin() bool {
	return c == CoinType_CoinType_DOGE
}

// IsSonr returns true if the coin type is sonr.
func (c CoinType) IsSonr() bool {
	return c == CoinType_CoinType_SONR
}

// DidMethod returns the DID method for the given coin type.
func (c CoinType) DidMethod() string {
	if c.IsBitcoin() {
		return "btcr"
	}
	if c.IsEthereum() {
		return "ethr"
	}
	if c.IsSonr() {
		return "sonr"
	}
	return strings.ToLower(c.Ticker())
}

// FormatAddress returns the address for the given public key for the spec of the coin type.
func (c CoinType) FormatAddress(pk *PubKey) string {
	if c.IsBitcoin() {
		return BitcoinAddress(pk)
	}
	if c.IsEthereum() {
		return EthereumAddress(pk)
	}
	if c.IsSonr() {
		addr, _ :=bech32.ConvertAndEncode("idx", pk.Address().Bytes())
		return addr
	}
	if c.IsCosmos() {
		addr, _ := bech32.ConvertAndEncode("cosmos", pk.Address().Bytes())
		return addr
	}
	return pk.Address().String()
}

// FormatDID returns the DID for the given public key for the spec of the coin type, along with the address.
func (c CoinType) FormatDID(pk *PubKey) (string, string) {
	return fmt.Sprintf("did:%s:%s", c.DidMethod(), c.FormatAddress(pk)), c.FormatAddress(pk)
}
