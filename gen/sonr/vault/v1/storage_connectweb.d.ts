// @generated by protoc-gen-connect-web v0.8.3
// @generated from file sonr/vault/v1/storage.proto (package sonrhq.sonr.vault.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

// Package Motor is used for defining a Motor node and its properties.

import { RefreshSharesRequest, RefreshSharesResponse } from "./storage_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * Vault is the service used for managing a node's keypair.
 *
 * @generated from service sonrhq.sonr.vault.v1.VaultStorage
 */
export declare const VaultStorage: {
  readonly typeName: "sonrhq.sonr.vault.v1.VaultStorage",
  readonly methods: {
    /**
     * Refresh Shares
     *
     * {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
     * It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
     *
     * #### {{.RequestType.Name}}
     * | Name | Type | Description |
     * | ---- | ---- | ----------- |{{range .RequestType.Fields}}
     * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
     *
     *
     * #### {{.ResponseType.Name}}
     * | Name | Type | Description |
     * | ---- | ---- | ----------- |{{range .ResponseType.Fields}}
     * | {{.Name}} | {{if eq .Label.String "LABEL_REPEATED"}}[]{{end}}{{.Type}} | {{fieldcomments .Message .}} | {{end}}
     *
     * @generated from rpc sonrhq.sonr.vault.v1.VaultStorage.RefreshShares
     */
    readonly refreshShares: {
      readonly name: "RefreshShares",
      readonly I: typeof RefreshSharesRequest,
      readonly O: typeof RefreshSharesResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

