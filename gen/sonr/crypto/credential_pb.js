// @generated by protoc-gen-es v1.1.0
// @generated from file sonr/crypto/credential.proto (package sonrhq.sonr.crypto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * WebauthnCredential contains all needed information about a WebAuthn credential for storage
 *
 * @generated from message sonrhq.sonr.crypto.WebauthnCredential
 */
export const WebauthnCredential = proto3.makeMessageType(
  "sonrhq.sonr.crypto.WebauthnCredential",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "public_key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "attestation_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "transport", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "authenticator", kind: "message", T: WebauthnAuthenticator },
  ],
);

/**
 * WebauthnAuthenticator contains certificate information about a WebAuthn authenticator
 *
 * @generated from message sonrhq.sonr.crypto.WebauthnAuthenticator
 */
export const WebauthnAuthenticator = proto3.makeMessageType(
  "sonrhq.sonr.crypto.WebauthnAuthenticator",
  () => [
    { no: 1, name: "aaguid", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "sign_count", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 3, name: "clone_warning", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

