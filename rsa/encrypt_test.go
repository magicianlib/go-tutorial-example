package rsa_example_test

import (
	"encoding/base64"
	"path/filepath"
	"testing"

	rsa_example "itknown.io/rsa"
)

func TestEncrypt(t *testing.T) {
	abs, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	der := filepath.Join(abs, "der", "id.der.pub")
	ciphertext, err := rsa_example.Encrypt(der, []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(base64.StdEncoding.EncodeToString(ciphertext))
	}

}

func TestDecrypt(t *testing.T) {
	abs, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	plaintext := []byte("hello world")

	pubDer := filepath.Join(abs, "der", "id.der.pub")
	ciphertext, err := rsa_example.Encrypt(pubDer, plaintext)
	if err != nil {
		t.Fatal(err)
	}

	priDer := filepath.Join(abs, "der", "id.der")
	decode, err := rsa_example.Decrypt(priDer, ciphertext)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(decode))

		if len(decode) != len(plaintext) {
			t.Fatal("解密失败")
		} else {
			for i := 0; i < len(plaintext); i++ {
				if v := plaintext[i] ^ decode[i]; v != 0 {
					t.Fatal("解密失败")
				}
			}
		}
	}
}
