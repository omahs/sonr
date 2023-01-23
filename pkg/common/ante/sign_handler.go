package ante

import (
	// cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

type CmpSignModeHandler struct {
}

func NewCmpSignModeHandler() authsigning.SignModeHandler {
	return CmpSignModeHandler{}
}

func (h CmpSignModeHandler) DefaultMode() signingtypes.SignMode {
	return signingtypes.SignMode_SIGN_MODE_DIRECT
}

func (h CmpSignModeHandler) Modes() []signingtypes.SignMode {
	return []signingtypes.SignMode{
		signingtypes.SignMode_SIGN_MODE_DIRECT,
		signingtypes.SignMode_SIGN_MODE_DIRECT_AUX,
		signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
	}
}

func (h CmpSignModeHandler) GetSignBytes(mode signing.SignMode, data authsigning.SignerData, tx sdk.Tx) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
