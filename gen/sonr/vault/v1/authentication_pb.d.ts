// @generated by protoc-gen-es v1.1.0
// @generated from file sonr/vault/v1/authentication.proto (package sonrhq.sonr.vault.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

// Package Motor is used for defining a Motor node and its properties.

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { DidDocument } from "../../../core/identity/did_pb.js";
import type { AccountInfo } from "../../common/info_pb.js";

/**
 * RegisterStartRequest is the request to register a new account.
 *
 * @generated from message sonrhq.sonr.vault.v1.RegisterStartRequest
 */
export declare class RegisterStartRequest extends Message<RegisterStartRequest> {
  /**
   * The origin of the request. This is used to query the Blockchain for the Service DID.
   *
   * @generated from field: string origin = 1;
   */
  origin: string;

  /**
   * The user defined label for the device.
   *
   * @generated from field: string device_label = 2;
   */
  deviceLabel: string;

  /**
   * The security threshold for the wallet account.
   *
   * @generated from field: int32 security_threshold = 3;
   */
  securityThreshold: number;

  /**
   * The recovery passcode for the wallet account.
   *
   * @generated from field: string passcode = 4;
   */
  passcode: string;

  /**
   * The Unique Identifier for the client device. Typically in a cookie.
   *
   * @generated from field: string uuid = 5;
   */
  uuid: string;

  constructor(data?: PartialMessage<RegisterStartRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "sonrhq.sonr.vault.v1.RegisterStartRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterStartRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterStartRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterStartRequest;

  static equals(a: RegisterStartRequest | PlainMessage<RegisterStartRequest> | undefined, b: RegisterStartRequest | PlainMessage<RegisterStartRequest> | undefined): boolean;
}

/**
 * RegisterStartResponse is the response to a Register request.
 *
 * @generated from message sonrhq.sonr.vault.v1.RegisterStartResponse
 */
export declare class RegisterStartResponse extends Message<RegisterStartResponse> {
  /**
   * Credential options for the user to sign with WebAuthn.
   *
   * @generated from field: string creation_options = 1;
   */
  creationOptions: string;

  /**
   * Relaying party id for the request.
   *
   * @generated from field: string rp_id = 2;
   */
  rpId: string;

  /**
   * Relaying party name for the request.
   *
   * @generated from field: string rp_name = 3;
   */
  rpName: string;

  constructor(data?: PartialMessage<RegisterStartResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "sonrhq.sonr.vault.v1.RegisterStartResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterStartResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterStartResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterStartResponse;

  static equals(a: RegisterStartResponse | PlainMessage<RegisterStartResponse> | undefined, b: RegisterStartResponse | PlainMessage<RegisterStartResponse> | undefined): boolean;
}

/**
 * RegisterFinishRequest is the request to CreateAccount a new key from the private key.
 *
 * @generated from message sonrhq.sonr.vault.v1.RegisterFinishRequest
 */
export declare class RegisterFinishRequest extends Message<RegisterFinishRequest> {
  /**
   * The previously generated session id.
   *
   * @generated from field: string uuid = 1;
   */
  uuid: string;

  /**
   * The signed credential response from the user.
   *
   * @generated from field: string credential_response = 2;
   */
  credentialResponse: string;

  /**
   * The origin of the request. This is used to query the Blockchain for the Service DID.
   *
   * @generated from field: string origin = 3;
   */
  origin: string;

  constructor(data?: PartialMessage<RegisterFinishRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "sonrhq.sonr.vault.v1.RegisterFinishRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterFinishRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterFinishRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterFinishRequest;

  static equals(a: RegisterFinishRequest | PlainMessage<RegisterFinishRequest> | undefined, b: RegisterFinishRequest | PlainMessage<RegisterFinishRequest> | undefined): boolean;
}

/**
 * RegisterFinishResponse is the response to a CreateAccount request.
 *
 * @generated from message sonrhq.sonr.vault.v1.RegisterFinishResponse
 */
export declare class RegisterFinishResponse extends Message<RegisterFinishResponse> {
  /**
   * The id of the account.
   *
   * @generated from field: bytes id = 1;
   */
  id: Uint8Array;

  /**
   * The address of the account.
   *
   * @generated from field: string address = 2;
   */
  address: string;

  /**
   * Relaying party id for the request.
   *
   * @generated from field: string rp_id = 3;
   */
  rpId: string;

  /**
   * Relaying party name for the request.
   *
   * @generated from field: string rp_name = 4;
   */
  rpName: string;

  /**
   * The DID Document for the wallet.
   *
   * @generated from field: sonrhq.core.identity.DidDocument did_document = 5;
   */
  didDocument?: DidDocument;

  /**
   * The account info for the wallet.
   *
   * @generated from field: sonrhq.sonr.common.AccountInfo account_info = 6;
   */
  accountInfo?: AccountInfo;

  /**
   * The UCAN token authorization header for subsequent requests.
   *
   * @generated from field: bytes ucan_token_header = 7;
   */
  ucanTokenHeader: Uint8Array;

  constructor(data?: PartialMessage<RegisterFinishResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "sonrhq.sonr.vault.v1.RegisterFinishResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterFinishResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterFinishResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterFinishResponse;

  static equals(a: RegisterFinishResponse | PlainMessage<RegisterFinishResponse> | undefined, b: RegisterFinishResponse | PlainMessage<RegisterFinishResponse> | undefined): boolean;
}

/**
 * LoginStartRequest is the request to login to an account.
 *
 * @generated from message sonrhq.sonr.vault.v1.LoginStartRequest
 */
export declare class LoginStartRequest extends Message<LoginStartRequest> {
  /**
   * The origin of the request. This is used to query the Blockchain for the Service DID.
   *
   * @generated from field: string origin = 1;
   */
  origin: string;

  /**
   * The Sonr account address for the user.
   *
   * @generated from field: string account_address = 2;
   */
  accountAddress: string;

  constructor(data?: PartialMessage<LoginStartRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "sonrhq.sonr.vault.v1.LoginStartRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginStartRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginStartRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginStartRequest;

  static equals(a: LoginStartRequest | PlainMessage<LoginStartRequest> | undefined, b: LoginStartRequest | PlainMessage<LoginStartRequest> | undefined): boolean;
}

/**
 * LoginStartResponse is the response to a Login request.
 *
 * @generated from message sonrhq.sonr.vault.v1.LoginStartResponse
 */
export declare class LoginStartResponse extends Message<LoginStartResponse> {
  /**
   * Success is true if the account exists.
   *
   * @generated from field: bool success = 1;
   */
  success: boolean;

  /**
   * The account address for the user.
   *
   * @generated from field: string account_address = 2;
   */
  accountAddress: string;

  /**
   * Json encoded WebAuthn credential options for the user to sign with.
   *
   * @generated from field: string credential_options = 3;
   */
  credentialOptions: string;

  /**
   * Relaying party id for the request.
   *
   * @generated from field: string rp_id = 4;
   */
  rpId: string;

  /**
   * Relaying party name for the request.
   *
   * @generated from field: string rp_name = 5;
   */
  rpName: string;

  constructor(data?: PartialMessage<LoginStartResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "sonrhq.sonr.vault.v1.LoginStartResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginStartResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginStartResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginStartResponse;

  static equals(a: LoginStartResponse | PlainMessage<LoginStartResponse> | undefined, b: LoginStartResponse | PlainMessage<LoginStartResponse> | undefined): boolean;
}

/**
 * LoginFinishRequest is the request to login to an account.
 *
 * @generated from message sonrhq.sonr.vault.v1.LoginFinishRequest
 */
export declare class LoginFinishRequest extends Message<LoginFinishRequest> {
  /**
   * Address of the account to login to.
   *
   * @generated from field: string account_address = 1;
   */
  accountAddress: string;

  /**
   * The signed credential response from the user.
   *
   * @generated from field: string credential_response = 2;
   */
  credentialResponse: string;

  /**
   * The origin of the request. This is used to query the Blockchain for the Service DID.
   *
   * @generated from field: string origin = 3;
   */
  origin: string;

  constructor(data?: PartialMessage<LoginFinishRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "sonrhq.sonr.vault.v1.LoginFinishRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginFinishRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginFinishRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginFinishRequest;

  static equals(a: LoginFinishRequest | PlainMessage<LoginFinishRequest> | undefined, b: LoginFinishRequest | PlainMessage<LoginFinishRequest> | undefined): boolean;
}

/**
 * LoginFinishResponse is the response to a Login request.
 *
 * @generated from message sonrhq.sonr.vault.v1.LoginFinishResponse
 */
export declare class LoginFinishResponse extends Message<LoginFinishResponse> {
  /**
   * Success is true if the account exists.
   *
   * @generated from field: bool success = 1;
   */
  success: boolean;

  /**
   * The account address for the user.
   *
   * @generated from field: string account_address = 2;
   */
  accountAddress: string;

  /**
   * Relaying party id for the request.
   *
   * @generated from field: string rp_id = 3;
   */
  rpId: string;

  /**
   * Relaying party name for the request.
   *
   * @generated from field: string rp_name = 4;
   */
  rpName: string;

  /**
   * The DID Document for the wallet.
   *
   * @generated from field: sonrhq.core.identity.DidDocument did_document = 5;
   */
  didDocument?: DidDocument;

  /**
   * The account info for the wallet.
   *
   * @generated from field: sonrhq.sonr.common.AccountInfo account_info = 6;
   */
  accountInfo?: AccountInfo;

  /**
   * The UCAN token authorization header for subsequent requests.
   *
   * @generated from field: bytes ucan_token_header = 7;
   */
  ucanTokenHeader: Uint8Array;

  constructor(data?: PartialMessage<LoginFinishResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "sonrhq.sonr.vault.v1.LoginFinishResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginFinishResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginFinishResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginFinishResponse;

  static equals(a: LoginFinishResponse | PlainMessage<LoginFinishResponse> | undefined, b: LoginFinishResponse | PlainMessage<LoginFinishResponse> | undefined): boolean;
}

