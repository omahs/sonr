// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: sonr/vault/v1/accounts.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/sonrhq/core/types/vault/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// VaultAccountsName is the fully-qualified name of the VaultAccounts service.
	VaultAccountsName = "sonrhq.sonr.vault.v1.VaultAccounts"
)

// VaultAccountsClient is a client for the sonrhq.sonr.vault.v1.VaultAccounts service.
type VaultAccountsClient interface {
	// Create a new account
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	CreateAccount(context.Context, *connect_go.Request[v1.CreateAccountRequest]) (*connect_go.Response[v1.CreateAccountResponse], error)
	// List the accounts
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	ListAccounts(context.Context, *connect_go.Request[v1.ListAccountsRequest]) (*connect_go.Response[v1.ListAccountsResponse], error)
	// Get Account
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	GetAccount(context.Context, *connect_go.Request[v1.GetAccountRequest]) (*connect_go.Response[v1.GetAccountResponse], error)
	// Delete Account
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	DeleteAccount(context.Context, *connect_go.Request[v1.DeleteAccountRequest]) (*connect_go.Response[v1.DeleteAccountResponse], error)
}

// NewVaultAccountsClient constructs a client for the sonrhq.sonr.vault.v1.VaultAccounts service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewVaultAccountsClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) VaultAccountsClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &vaultAccountsClient{
		createAccount: connect_go.NewClient[v1.CreateAccountRequest, v1.CreateAccountResponse](
			httpClient,
			baseURL+"/sonrhq.sonr.vault.v1.VaultAccounts/CreateAccount",
			opts...,
		),
		listAccounts: connect_go.NewClient[v1.ListAccountsRequest, v1.ListAccountsResponse](
			httpClient,
			baseURL+"/sonrhq.sonr.vault.v1.VaultAccounts/ListAccounts",
			opts...,
		),
		getAccount: connect_go.NewClient[v1.GetAccountRequest, v1.GetAccountResponse](
			httpClient,
			baseURL+"/sonrhq.sonr.vault.v1.VaultAccounts/GetAccount",
			opts...,
		),
		deleteAccount: connect_go.NewClient[v1.DeleteAccountRequest, v1.DeleteAccountResponse](
			httpClient,
			baseURL+"/sonrhq.sonr.vault.v1.VaultAccounts/DeleteAccount",
			opts...,
		),
	}
}

// vaultAccountsClient implements VaultAccountsClient.
type vaultAccountsClient struct {
	createAccount *connect_go.Client[v1.CreateAccountRequest, v1.CreateAccountResponse]
	listAccounts  *connect_go.Client[v1.ListAccountsRequest, v1.ListAccountsResponse]
	getAccount    *connect_go.Client[v1.GetAccountRequest, v1.GetAccountResponse]
	deleteAccount *connect_go.Client[v1.DeleteAccountRequest, v1.DeleteAccountResponse]
}

// CreateAccount calls sonrhq.sonr.vault.v1.VaultAccounts.CreateAccount.
func (c *vaultAccountsClient) CreateAccount(ctx context.Context, req *connect_go.Request[v1.CreateAccountRequest]) (*connect_go.Response[v1.CreateAccountResponse], error) {
	return c.createAccount.CallUnary(ctx, req)
}

// ListAccounts calls sonrhq.sonr.vault.v1.VaultAccounts.ListAccounts.
func (c *vaultAccountsClient) ListAccounts(ctx context.Context, req *connect_go.Request[v1.ListAccountsRequest]) (*connect_go.Response[v1.ListAccountsResponse], error) {
	return c.listAccounts.CallUnary(ctx, req)
}

// GetAccount calls sonrhq.sonr.vault.v1.VaultAccounts.GetAccount.
func (c *vaultAccountsClient) GetAccount(ctx context.Context, req *connect_go.Request[v1.GetAccountRequest]) (*connect_go.Response[v1.GetAccountResponse], error) {
	return c.getAccount.CallUnary(ctx, req)
}

// DeleteAccount calls sonrhq.sonr.vault.v1.VaultAccounts.DeleteAccount.
func (c *vaultAccountsClient) DeleteAccount(ctx context.Context, req *connect_go.Request[v1.DeleteAccountRequest]) (*connect_go.Response[v1.DeleteAccountResponse], error) {
	return c.deleteAccount.CallUnary(ctx, req)
}

// VaultAccountsHandler is an implementation of the sonrhq.sonr.vault.v1.VaultAccounts service.
type VaultAccountsHandler interface {
	// Create a new account
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	CreateAccount(context.Context, *connect_go.Request[v1.CreateAccountRequest]) (*connect_go.Response[v1.CreateAccountResponse], error)
	// List the accounts
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	ListAccounts(context.Context, *connect_go.Request[v1.ListAccountsRequest]) (*connect_go.Response[v1.ListAccountsResponse], error)
	// Get Account
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	GetAccount(context.Context, *connect_go.Request[v1.GetAccountRequest]) (*connect_go.Response[v1.GetAccountResponse], error)
	// Delete Account
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	DeleteAccount(context.Context, *connect_go.Request[v1.DeleteAccountRequest]) (*connect_go.Response[v1.DeleteAccountResponse], error)
}

// NewVaultAccountsHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewVaultAccountsHandler(svc VaultAccountsHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/sonrhq.sonr.vault.v1.VaultAccounts/CreateAccount", connect_go.NewUnaryHandler(
		"/sonrhq.sonr.vault.v1.VaultAccounts/CreateAccount",
		svc.CreateAccount,
		opts...,
	))
	mux.Handle("/sonrhq.sonr.vault.v1.VaultAccounts/ListAccounts", connect_go.NewUnaryHandler(
		"/sonrhq.sonr.vault.v1.VaultAccounts/ListAccounts",
		svc.ListAccounts,
		opts...,
	))
	mux.Handle("/sonrhq.sonr.vault.v1.VaultAccounts/GetAccount", connect_go.NewUnaryHandler(
		"/sonrhq.sonr.vault.v1.VaultAccounts/GetAccount",
		svc.GetAccount,
		opts...,
	))
	mux.Handle("/sonrhq.sonr.vault.v1.VaultAccounts/DeleteAccount", connect_go.NewUnaryHandler(
		"/sonrhq.sonr.vault.v1.VaultAccounts/DeleteAccount",
		svc.DeleteAccount,
		opts...,
	))
	return "/sonrhq.sonr.vault.v1.VaultAccounts/", mux
}

// UnimplementedVaultAccountsHandler returns CodeUnimplemented from all methods.
type UnimplementedVaultAccountsHandler struct{}

func (UnimplementedVaultAccountsHandler) CreateAccount(context.Context, *connect_go.Request[v1.CreateAccountRequest]) (*connect_go.Response[v1.CreateAccountResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.sonr.vault.v1.VaultAccounts.CreateAccount is not implemented"))
}

func (UnimplementedVaultAccountsHandler) ListAccounts(context.Context, *connect_go.Request[v1.ListAccountsRequest]) (*connect_go.Response[v1.ListAccountsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.sonr.vault.v1.VaultAccounts.ListAccounts is not implemented"))
}

func (UnimplementedVaultAccountsHandler) GetAccount(context.Context, *connect_go.Request[v1.GetAccountRequest]) (*connect_go.Response[v1.GetAccountResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.sonr.vault.v1.VaultAccounts.GetAccount is not implemented"))
}

func (UnimplementedVaultAccountsHandler) DeleteAccount(context.Context, *connect_go.Request[v1.DeleteAccountRequest]) (*connect_go.Response[v1.DeleteAccountResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.sonr.vault.v1.VaultAccounts.DeleteAccount is not implemented"))
}
