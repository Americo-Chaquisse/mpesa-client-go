package mpesa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

// AuthToken to be used as bearer token invoking mpesa api
// ref: https://developer.mpesa.vm.co.mz/documentation/
func AuthToken(publicKeyString string, apiKey string) (string, error) {
	if len(publicKeyString) == 0 {
		return "", errors.New("public key is not defined")
	}
	if len(apiKey) == 0 {
		return "", errors.New("api key is not defined")
	}
	// Generate a decoded Base64 string from the Public Key
	decodedString, err := base64.StdEncoding.DecodeString(publicKeyString)
	if err != nil {
		return "", err
	}

	// Generate an instance of an RSA cipher and use the decoded Base64 string as the input
	var publicKey *rsa.PublicKey
	if unparsed, err := x509.ParsePKIXPublicKey(decodedString); err != nil {
		return "", err
	} else {
		publicKey = unparsed.(*rsa.PublicKey)
	}

	// Encode the API Key with the RSA cipher
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(apiKey))
	if err != nil {
		return "", err
	}

	// Digest as Base64 string format
	return base64.StdEncoding.EncodeToString(cipher), nil
}
