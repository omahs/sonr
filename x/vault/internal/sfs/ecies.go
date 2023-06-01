package sfs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	fmt "fmt"
	io "io"
	"math/big"

	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/sonrhq/core/x/vault/types"
	"golang.org/x/crypto/hkdf"
)

// This function takes a public key and plaintext as input. It first converts the public key from webauthncose.EC2PublicKeyData format
// to ecdsa.PublicKey format. It then generates an ephemeral key pair, derives a shared secret, and derives encryption and MAC keys
// from the shared secret. The plaintext is then encrypted with AES-256-GCM and a MAC tag is computed. The function finally encodes the
// ephemeral public key and concatenates the public key, IV, ciphertext, and MAC tag into a single byte slice.
func eciesEncryptFromPub(publicKeyData webauthncose.EC2PublicKeyData, plaintext []byte) ([]byte, error) {
	// Convert the X and Y coordinates of the public key to big.Int values
	x := new(big.Int).SetBytes(publicKeyData.XCoord)
	y := new(big.Int).SetBytes(publicKeyData.YCoord)

	// Create an ECDSA public key from the X and Y coordinates and the curve identifier
	curve := getCurve(publicKeyData.Algorithm)
	if curve == nil {
		return nil, fmt.Errorf("unsupported curve identifier: %d", publicKeyData.Algorithm)
	}
	publicKey := ecdsa.PublicKey{Curve: curve, X: x, Y: y}
	// Generate a random key pair for the ephemeral keys
	ephemeralPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ephemeral private key: %w", err)
	}

	// Derive the shared secret
	sharedSecretX, _ := publicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, ephemeralPrivateKey.D.Bytes())
	sharedSecret := sharedSecretX.Bytes()

	// Derive the encryption key and MAC key from the shared secret
	encryptionKey, macKey := deriveKeys(sharedSecret)

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}

	// Encrypt the plaintext using AES-256-GCM with the encryption key and IV
	ciphertext := make([]byte, len(plaintext))
	c, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}
	gcm, err := cipher.NewGCMWithNonceSize(c, aes.BlockSize)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}
	gcm.Seal(ciphertext[:0], iv, plaintext, nil)

	// Compute the MAC tag
	mac := hmac.New(sha256.New, macKey)
	mac.Write(iv)
	mac.Write(ciphertext)
	tag := mac.Sum(nil)

	// Encode the ephemeral public key
	ephemeralPublicKeyBytes := elliptic.Marshal(elliptic.P256(), ephemeralPrivateKey.PublicKey.X, ephemeralPrivateKey.PublicKey.Y)
	ephemeralPublicKey, err := x509.ParsePKIXPublicKey(ephemeralPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to create ephemeral public key: %w", err)
	}
	encodedPublicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(ephemeralPublicKey.(*rsa.PublicKey)),
	})

	// Concatenate the public key, IV, ciphertext, and MAC tag
	result := make([]byte, len(encodedPublicKey)+len(iv)+len(ciphertext)+len(tag))
	copy(result[:len(encodedPublicKey)], encodedPublicKey)
	copy(result[len(encodedPublicKey):len(encodedPublicKey)+len(iv)], iv)
	copy(result[len(encodedPublicKey)+len(iv):len(encodedPublicKey)+len(iv)+len(ciphertext)], ciphertext)
	copy(result[len(encodedPublicKey)+len(iv)+len(ciphertext):], tag)
	return result, nil
}

// This function derives encryption and MAC keys from the shared secret using the HKDF
// (HMAC-based Extract-and-Expand Key Derivation Function) function.
func deriveKeys(sharedSecret []byte) ([]byte, []byte) {
	// Use HKDF to derive the encryption key and MAC key from the shared secret
	info := []byte("encryption key")
	encryptionKey := hkdf.New(sha256.New, sharedSecret, nil, info)
	info = []byte("MAC key")
	macKey := hkdf.New(sha256.New, sharedSecret, nil, info)
	encKeyBytes := make([]byte, 32)
	if _, err := io.ReadFull(encryptionKey, encKeyBytes); err != nil {
		return nil, nil
	}
	macKeyBytes := make([]byte, 32)
	if _, err := io.ReadFull(macKey, macKeyBytes); err != nil {
		return nil, nil
	}
	return encKeyBytes, macKeyBytes
}

// This function derives a private key from a WebAuthn credential. It parses the public key from the credential, generates an ephemeral
// private key, derives a shared secret using ECDH (Elliptic Curve Diffie-Hellman), and then derives a 256-bit key from the shared secret
// using HKDF. The derived key is then used to create a new private key.
func derivePrivateKey(publicKey []byte) (*ecdsa.PrivateKey, error) {
	// Parse the public key from the credential
	pubKeyFace, err := webauthncose.ParsePublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	pubKey, ok := pubKeyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, fmt.Errorf("public key is not an EC2 key")
	}

	// Convert the x and y coordinates to *big.Int values
	xCoord := new(big.Int).SetBytes(pubKey.XCoord)
	yCoord := new(big.Int).SetBytes(pubKey.YCoord)

	// Generate a new ephemeral private key
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ephemeral private key: %w", err)
	}

	// Derive the shared secret using ECDH
	sharedX, _ := privKey.Curve.ScalarMult(xCoord, yCoord, privKey.D.Bytes())
	sharedSecret := sharedX.Bytes()

	// Derive a 256-bit key using HKDF
	keyBytes := make([]byte, 32)
	info := []byte("webauthn-secret")
	hkdf := hkdf.New(sha256.New, sharedSecret, nil, info)
	if _, err := io.ReadFull(hkdf, keyBytes); err != nil {
		return nil, fmt.Errorf("failed to derive key: %w", err)
	}

	// Create a new private key from the derived key
	return &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     xCoord,
			Y:     yCoord,
		},
		D: new(big.Int).SetBytes(keyBytes),
	}, nil
}

// This function calculates the shared secret between a private key and a public key using ECDH.
func sharedSecret(privateKey *ecdsa.PrivateKey, publicKey webauthncose.EC2PublicKeyData) ([]byte, error) {
	// Convert the X and Y coordinates of the public key to big.Int values
	x := new(big.Int).SetBytes(publicKey.XCoord)
	y := new(big.Int).SetBytes(publicKey.YCoord)

	// Create an ECDSA public key from the X and Y coordinates and the curve identifier
	curve := getCurve(publicKey.Algorithm)
	if curve == nil {
		return nil, fmt.Errorf("unsupported curve identifier: %d", publicKey.Algorithm)
	}
	publicKeyEcdsa := ecdsa.PublicKey{Curve: curve, X: x, Y: y}

	// Calculate the shared secret using ECDH
	x, _ = curve.ScalarMult(publicKeyEcdsa.X, publicKeyEcdsa.Y, privateKey.D.Bytes())
	return x.Bytes(), nil
}

// This function determines the elliptic curve to be used based on the COSE algorithm identifier in the public key data.
func getCurve(curveID int64) elliptic.Curve {
	var curve elliptic.Curve
	switch webauthncose.COSEAlgorithmIdentifier(curveID) {
	case webauthncose.AlgES512: // IANA COSE code for ECDSA w/ SHA-512
		curve = elliptic.P521()
	case webauthncose.AlgES384: // IANA COSE code for ECDSA w/ SHA-384
		curve = elliptic.P384()
	case webauthncose.AlgES256: // IANA COSE code for ECDSA w/ SHA-256
		curve = elliptic.P256()
	default:
		return nil
	}
	return curve
}

// encryptEcies is used to encrypt a message for the credential
func encryptEcies(credential *servicetypes.WebauthnCredential, data []byte) ([]byte, error) {
	// Get the public key from the credential
	keyFace, err := webauthncose.ParsePublicKey(credential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, fmt.Errorf("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(credential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key: %w", err)
	}
	sharedSecret, err := sharedSecret(privateKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive shared secret: %w", err)
	}
	// Use the shared secret as the encryption key
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}

	// Encrypt the data using AES-GCM
	ciphertext := make([]byte, len(data))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}
	gcm.Seal(ciphertext[:0], iv, data, nil)

	// Encrypt the AES-GCM key using ECIES
	encryptedKey, err := eciesEncryptFromPub(publicKey, sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt key: %w", err)
	}

	// Concatenate the IV and ciphertext into a single byte slice
	result := make([]byte, len(iv)+len(ciphertext)+len(encryptedKey))
	copy(result[:len(iv)], iv)
	copy(result[len(iv):len(iv)+len(ciphertext)], ciphertext)
	copy(result[len(iv)+len(ciphertext):], encryptedKey)

	return result, nil
}

// decryptEcies is used to decrypt a message for the credential
func decryptEcies(credential *servicetypes.WebauthnCredential, data []byte) ([]byte, error) {
	// Get the public key from the credential
	keyFace, err := webauthncose.ParsePublicKey(credential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, fmt.Errorf("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(credential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key: %w", err)
	}

	// Derive the shared secret using ECDH and the WebAuthn credential
	sharedSecret, err := sharedSecret(privateKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive shared secret: %w", err)
	}

	// Use the shared secret as the decryption key
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Split the IV and ciphertext from the encrypted data
	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	// Decrypt the ciphertext using AES-GCM
	plaintext := make([]byte, len(ciphertext))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}
	if _, err := gcm.Open(plaintext[:0], iv, ciphertext, nil); err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %w", err)
	}
	return plaintext, nil
}

// insertECIESKeyshare inserts a raw keyshare encrypted with ECIES derived from a WebAuthn credential
func insertECIESKeyshare(ks types.KeyShare, credential *servicetypes.WebauthnCredential) error {
	dat := ks.Bytes()
	encDat, err := encryptEcies(credential, dat)
	if err != nil {
		return fmt.Errorf("failed to encrypt keyshare: %w", err)
	}
	_, err = ksTable.Put(ctx, types.KeysharePrefix(ks.Did()), encDat)
	if err != nil {
		return fmt.Errorf("failed to put keyshare: %w", err)
	}
	return nil
}

// This function returns a key share and an error for a given DID and webauthn credential.
func getECIESKeyshare(ksDid string, credential *servicetypes.WebauthnCredential) (types.KeyShare, error) {
	ksr, err := types.ParseKeyShareDID(ksDid)
	if err != nil {
		return nil, fmt.Errorf("failed to parse keyshare DID: %w", err)
	}
	vEnc, err := ksTable.Get(ctx, types.KeysharePrefix(ksDid))
	if err != nil {
		return nil, fmt.Errorf("failed to get keyshare: %w", err)
	}
	vBiz, err := decryptEcies(credential, vEnc)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt keyshare: %w", err)
	}
	ks, err := types.NewKeyshare(ksDid, vBiz, ksr.CoinType)
	if err != nil {
		return nil, fmt.Errorf("failed to parse keyshare: %w", err)
	}
	return ks, nil
}
