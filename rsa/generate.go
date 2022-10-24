package rsa_example

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// GenerateRsaKeyPair generate rsa key-pair
func GenerateRsaKeyPair(bits int, writeFilename string) (privDer, pubDer []byte, err error) {
	keyPair, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	privDer, pubDer = _x509Der(keyPair)

	if writeFilename != "" {

		if err := os.WriteFile(writeFilename, privDer, 0700); err != nil {
			return nil, nil, err
		}

		if err := os.WriteFile(writeFilename+".pub", pubDer, 0755); err != nil {
			return nil, nil, err
		}
	}

	return privDer, pubDer, nil
}

// _x509Der transform rsa key-pair to x509 format
func _x509Der(keyPair *rsa.PrivateKey) (privDer, pubDer []byte) {

	// DER 编码数据

	privDer = pem.EncodeToMemory(&pem.Block{
		Type:    "DER(x509) PUBLIC KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(keyPair),
	})

	pubDer = pem.EncodeToMemory(&pem.Block{
		Type:    "DER(x509) PUBLIC KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PublicKey(&keyPair.PublicKey),
	})

	return privDer, pubDer
}
