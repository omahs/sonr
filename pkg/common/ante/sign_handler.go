package ante

import (
	// cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// CmpSignModeHandler Is a struct that implements the `SignModeHandler` interface.
type CmpSignModeHandler struct{}

// NewCmpSignModeHandler returns a new SignModeHandler which supports CMP based signing
func NewCmpSignModeHandler() authsigning.SignModeHandler {
	return CmpSignModeHandler{}
}

// DefaultMode is the default mode that is to be used with this handler if no
func (h CmpSignModeHandler) DefaultMode() signingtypes.SignMode {
	return signingtypes.SignMode_SIGN_MODE_DIRECT_AUX
}

// Modes is the list of modes supporting by this handler
func (h CmpSignModeHandler) Modes() []signingtypes.SignMode {
	return []signingtypes.SignMode{
		signingtypes.SignMode_SIGN_MODE_DIRECT,
		signingtypes.SignMode_SIGN_MODE_DIRECT_AUX,
		signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
	}
}

// GetSignBytes returns the sign bytes for the provided SignMode, SignerData and Tx,
func (h CmpSignModeHandler) GetSignBytes(mode signing.SignMode, data authsigning.SignerData, tx sdk.Tx) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
