package controller

import (
	"fmt"
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
)

type KeyShareParseResult struct {
	CoinType     crypto.CoinType
	AccountName  string
	KeyShareName string
}

// ParseKeyShareDid parses a keyshare DID into its components. The DID format is:
// did:{coin_type}:{account_address}#ks-{keyshare_name}
func ParseKeyShareDid(name string) (*KeyShareParseResult, error) {
	// Parse the DID
	parts := strings.Split(name, ":")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid keyshare DID: %s", name)
	}

	// Parse the coin type
	ct := crypto.CoinTypeFromDidMethod(parts[1])

	// Split the account address and keyshare name
	parts = strings.Split(parts[2], "#")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid keyshare DID: %s", name)
	}

	// Parse the account address
	accountName := parts[0]

	// Parse the keyshare name
	parts = strings.Split(parts[1], "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid keyshare DID: %s", name)
	}

	// Parse the keyshare name
	keyShareName := parts[1]



	return &KeyShareParseResult{
		CoinType:     ct,
		AccountName:  accountName,
		KeyShareName: keyShareName,
	}, nil
}
