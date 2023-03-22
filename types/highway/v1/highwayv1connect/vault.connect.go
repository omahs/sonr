// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: highway/v1/vault.proto

package highwayv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/sonrhq/core/types/highway/v1"
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
	// VaultName is the fully-qualified name of the Vault service.
	VaultName = "sonrhq.highway.v1.Vault"
)

// VaultClient is a client for the sonrhq.highway.v1.Vault service.
type VaultClient interface {
	// Assign Key
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Assign(context.Context, *connect_go.Request[v1.AssignKeyRequest]) (*connect_go.Response[v1.AssignKeyResponse], error)
	// Login Start
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Upload(context.Context, *connect_go.Request[v1.UploadShareRequest]) (*connect_go.Response[v1.UploadShareResponse], error)
	// Login Finish
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Sync(context.Context, *connect_go.Request[v1.SyncShareRequest]) (*connect_go.Response[v1.SyncShareResponse], error)
	// Register Start
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Refresh(context.Context, *connect_go.Request[v1.RefreshShareRequest]) (*connect_go.Response[v1.RefreshShareResponse], error)
}

// NewVaultClient constructs a client for the sonrhq.highway.v1.Vault service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewVaultClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) VaultClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &vaultClient{
		assign: connect_go.NewClient[v1.AssignKeyRequest, v1.AssignKeyResponse](
			httpClient,
			baseURL+"/sonrhq.highway.v1.Vault/Assign",
			opts...,
		),
		upload: connect_go.NewClient[v1.UploadShareRequest, v1.UploadShareResponse](
			httpClient,
			baseURL+"/sonrhq.highway.v1.Vault/Upload",
			opts...,
		),
		sync: connect_go.NewClient[v1.SyncShareRequest, v1.SyncShareResponse](
			httpClient,
			baseURL+"/sonrhq.highway.v1.Vault/Sync",
			opts...,
		),
		refresh: connect_go.NewClient[v1.RefreshShareRequest, v1.RefreshShareResponse](
			httpClient,
			baseURL+"/sonrhq.highway.v1.Vault/Refresh",
			opts...,
		),
	}
}

// vaultClient implements VaultClient.
type vaultClient struct {
	assign  *connect_go.Client[v1.AssignKeyRequest, v1.AssignKeyResponse]
	upload  *connect_go.Client[v1.UploadShareRequest, v1.UploadShareResponse]
	sync    *connect_go.Client[v1.SyncShareRequest, v1.SyncShareResponse]
	refresh *connect_go.Client[v1.RefreshShareRequest, v1.RefreshShareResponse]
}

// Assign calls sonrhq.highway.v1.Vault.Assign.
func (c *vaultClient) Assign(ctx context.Context, req *connect_go.Request[v1.AssignKeyRequest]) (*connect_go.Response[v1.AssignKeyResponse], error) {
	return c.assign.CallUnary(ctx, req)
}

// Upload calls sonrhq.highway.v1.Vault.Upload.
func (c *vaultClient) Upload(ctx context.Context, req *connect_go.Request[v1.UploadShareRequest]) (*connect_go.Response[v1.UploadShareResponse], error) {
	return c.upload.CallUnary(ctx, req)
}

// Sync calls sonrhq.highway.v1.Vault.Sync.
func (c *vaultClient) Sync(ctx context.Context, req *connect_go.Request[v1.SyncShareRequest]) (*connect_go.Response[v1.SyncShareResponse], error) {
	return c.sync.CallUnary(ctx, req)
}

// Refresh calls sonrhq.highway.v1.Vault.Refresh.
func (c *vaultClient) Refresh(ctx context.Context, req *connect_go.Request[v1.RefreshShareRequest]) (*connect_go.Response[v1.RefreshShareResponse], error) {
	return c.refresh.CallUnary(ctx, req)
}

// VaultHandler is an implementation of the sonrhq.highway.v1.Vault service.
type VaultHandler interface {
	// Assign Key
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Assign(context.Context, *connect_go.Request[v1.AssignKeyRequest]) (*connect_go.Response[v1.AssignKeyResponse], error)
	// Login Start
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Upload(context.Context, *connect_go.Request[v1.UploadShareRequest]) (*connect_go.Response[v1.UploadShareResponse], error)
	// Login Finish
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Sync(context.Context, *connect_go.Request[v1.SyncShareRequest]) (*connect_go.Response[v1.SyncShareResponse], error)
	// Register Start
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	// #### {{.RequestType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .RequestType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	//
	// #### {{.ResponseType.Name}}
	// | Name | Type | Description |
	// | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
	// | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
	Refresh(context.Context, *connect_go.Request[v1.RefreshShareRequest]) (*connect_go.Response[v1.RefreshShareResponse], error)
}

// NewVaultHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewVaultHandler(svc VaultHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/sonrhq.highway.v1.Vault/Assign", connect_go.NewUnaryHandler(
		"/sonrhq.highway.v1.Vault/Assign",
		svc.Assign,
		opts...,
	))
	mux.Handle("/sonrhq.highway.v1.Vault/Upload", connect_go.NewUnaryHandler(
		"/sonrhq.highway.v1.Vault/Upload",
		svc.Upload,
		opts...,
	))
	mux.Handle("/sonrhq.highway.v1.Vault/Sync", connect_go.NewUnaryHandler(
		"/sonrhq.highway.v1.Vault/Sync",
		svc.Sync,
		opts...,
	))
	mux.Handle("/sonrhq.highway.v1.Vault/Refresh", connect_go.NewUnaryHandler(
		"/sonrhq.highway.v1.Vault/Refresh",
		svc.Refresh,
		opts...,
	))
	return "/sonrhq.highway.v1.Vault/", mux
}

// UnimplementedVaultHandler returns CodeUnimplemented from all methods.
type UnimplementedVaultHandler struct{}

func (UnimplementedVaultHandler) Assign(context.Context, *connect_go.Request[v1.AssignKeyRequest]) (*connect_go.Response[v1.AssignKeyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.highway.v1.Vault.Assign is not implemented"))
}

func (UnimplementedVaultHandler) Upload(context.Context, *connect_go.Request[v1.UploadShareRequest]) (*connect_go.Response[v1.UploadShareResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.highway.v1.Vault.Upload is not implemented"))
}

func (UnimplementedVaultHandler) Sync(context.Context, *connect_go.Request[v1.SyncShareRequest]) (*connect_go.Response[v1.SyncShareResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.highway.v1.Vault.Sync is not implemented"))
}

func (UnimplementedVaultHandler) Refresh(context.Context, *connect_go.Request[v1.RefreshShareRequest]) (*connect_go.Response[v1.RefreshShareResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.highway.v1.Vault.Refresh is not implemented"))
}
