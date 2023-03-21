package highway

import (
	highwayv1connect "github.com/sonrhq/core/types/highway/v1/highwayv1connect"
)

type Highway struct {
	Authentication highwayv1connect.AuthenticationHandler
	Accounts 	 highwayv1connect.MpcHandler
	Vault 		 highwayv1connect.VaultHandler
}
