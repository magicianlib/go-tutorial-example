package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	priPEM, pubPEM, err := GenerateRsaKeyPair(4096, "")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(priPEM))
	fmt.Println(string(pubPEM))
}

func GenerateRsaKeyPair(bits int, filenamePrefix string) (priPEM, pubPEM []byte, err error) {
	keyPair, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	priPEM, pubPEM = _x509(keyPair)

	if filenamePrefix != "" {

		if err := os.WriteFile(filenamePrefix+".rsa", priPEM, 0700); err != nil {
			return nil, nil, err
		}

		if err := os.WriteFile(filenamePrefix+".rsa.pub", pubPEM, 0755); err != nil {
			return nil, nil, err
		}
	}

	return priPEM, pubPEM, nil
}

func _x509(keyPair *rsa.PrivateKey) (priPEM, pubPEM []byte) {
	priPEM = pem.EncodeToMemory(&pem.Block{
		Type:    "RSA(x509) PUBLIC KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(keyPair),
	})

	pubPEM = pem.EncodeToMemory(&pem.Block{
		Type:    "RSA(x509) PUBLIC KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PublicKey(&keyPair.PublicKey),
	})

	return priPEM, pubPEM
}
