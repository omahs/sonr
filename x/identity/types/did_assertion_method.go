// Utility functions for DID Assertion Method - https://w3c.github.io/did-core/#assertion
// I.e. Verification Material for Wallets. This is the default Verification Method for DID Documents. (snr, btc, eth, etc.)
package types

// KnownWalletPrefixes is an enum of known wallet prefixes
type ChainWalletPrefix int

const (
	// KnownWalletPrefixes is an enum of known wallet prefixes
	ChainWalletPrefixNone ChainWalletPrefix = iota
	ChainWalletPrefixSNR
	ChainWalletPrefixBTC
	ChainWalletPrefixETH
)

func NewWalletPrefix(prefix string) ChainWalletPrefix {
	switch prefix {
	case "snr":
		return ChainWalletPrefixSNR
	case "btc":
		return ChainWalletPrefixBTC
	case "eth":
		return ChainWalletPrefixETH
	case "0x":
		return ChainWalletPrefixETH
	default:
		return ChainWalletPrefixNone
	}
}

func (k ChainWalletPrefix) String() string {
	return [...]string{"account", "snr", "btc", "eth"}[k]
}

// Prefix returns the prefix of the wallet
func (k ChainWalletPrefix) Prefix() string {
	if k == ChainWalletPrefixETH {
		return "0x"
	}
	return k.String()
}

// AssertionMethodCount returns the number of Assertion Methods
func (vm *DidDocument) AssertionMethodCount() int {
	return vm.AssertionMethod.Count()
}

// FindAssertionMethod finds a VerificationMethod by its ID
func (d *DidDocument) FindAssertionMethod(id string) *VerificationMethod {
	return d.AssertionMethod.FindByID(id)
}

// FindAssertionMethodByFragment finds a VerificationMethod by its fragment
func (d *DidDocument) FindAssertionMethodByFragment(fragment string) *VerificationMethod {
	return d.AssertionMethod.FindByFragment(fragment)[0]
}

// AddAssertionMethod adds a VerificationMethod as AssertionMethod
// If the controller is not set, it will be set to the documents ID
func (d *DidDocument) AddAssertion(v *VerificationMethod) {
	if v.Controller == "" {
		v.Controller = d.ID
	}
	d.VerificationMethod.Add(v)
	d.AssertionMethod.Add(v)
}

// GetBlockchainAccountCount returns the number of Blockchain Accounts by the address prefix
func (d *DidDocument) GetBlockchainAccountCount(prefix string) int {
	return len(d.AssertionMethod.FindByFragment(prefix))
}

// ListBlockchainAccounts returns a list of Blockchain Accounts by the address prefix
func (d *DidDocument) ListBlockchainAccounts() []*VerificationMethod {
	accs := make([]*VerificationMethod, 0)
	for _, vm := range d.AssertionMethod.Data {
		if vm.VerificationMethod.IsBlockchainAccount() {
			accs = append(accs, vm.VerificationMethod)
		}
	}
	return accs
}
