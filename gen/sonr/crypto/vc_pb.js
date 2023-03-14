// @generated by protoc-gen-es v1.1.0
// @generated from file sonr/crypto/vc.proto (package sonrhq.sonr.crypto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { ProofType } from "./ssi_pb.js";

/**
 * Proof represents a credential/presentation proof as defined by the Linked Data Proofs 1.0 specification (https://w3c-ccg.github.io/ld-proofs/).
 *
 * @generated from message sonrhq.sonr.crypto.Proof
 */
export const Proof = proto3.makeMessageType(
  "sonrhq.sonr.crypto.Proof",
  () => [
    { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(ProofType) },
    { no: 2, name: "proof_purpose", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "verification_method", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "created", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "domain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * JSONWebSignature2020Proof is a VC proof with a signature according to JsonWebSignature2020
 *
 * @generated from message sonrhq.sonr.crypto.JSONWebSignature2020Proof
 */
export const JSONWebSignature2020Proof = proto3.makeMessageType(
  "sonrhq.sonr.crypto.JSONWebSignature2020Proof",
  () => [
    { no: 1, name: "proof", kind: "message", T: Proof },
    { no: 2, name: "jws", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * VerifiableCredential represents a credential as defined by the Verifiable Credentials Data Model 1.0 specification (https://www.w3.org/TR/vc-data-model/).
 *
 * @generated from message sonrhq.sonr.crypto.VerifiableCredential
 */
export const VerifiableCredential = proto3.makeMessageType(
  "sonrhq.sonr.crypto.VerifiableCredential",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "context", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "issuer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "issuance_date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "expiration_date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "credential_subject", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 8, name: "proof", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ],
);

