package chain

type ClientOriginAPI string

const (
	// List of known origin api endpoints.
	SonrLocalRpcOrigin  ClientOriginAPI = "localhost:9090"
	SonrPublicRpcOrigin ClientOriginAPI = "159.65.236.177:9090"
)
