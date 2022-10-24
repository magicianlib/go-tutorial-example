package rsa_example_test

import (
	"crypto"
	"path/filepath"
	"testing"

	rsa_example "itknown.io/rsa"
)

func TestSha256WithRsaSign(t *testing.T) {

	abs, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	data := []byte("你好, world")

	privDer := filepath.Join(abs, "der", "id.der")

	base64sign, err := rsa_example.HashWithRsaSign(privDer, crypto.SHA256, data)
	if err != nil {
		t.Fatal("Sha256WithRsaSign err", err)
	} else {
		t.Log("Sha256WithRsaSign:", base64sign)
	}
}

func TestMd5WithRsaSign(t *testing.T) {

	abs, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	data := []byte("你好, world")

	privDer := filepath.Join(abs, "der", "id.der")

	base64sign, err := rsa_example.HashWithRsaSign(privDer, crypto.MD5, data)
	if err != nil {
		t.Fatal("Md5WithRsaSign fail", err)
	} else {
		t.Log("Md5WithRsaSign:", base64sign)
	}
}

func TestSha256WithRsaSignVerify(t *testing.T) {
	abs, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	data := []byte("你好, world")

	privDer := filepath.Join(abs, "der", "id.der")
	pubDer := filepath.Join(abs, "der", "id.der.pub")

	base64sign, err := rsa_example.HashWithRsaSign(privDer, crypto.SHA256, data)
	if err != nil {
		t.Fatal("Sha256WithRsaSign err", err)
	} else {
		err := rsa_example.HashWithRsaSignVerify(pubDer, crypto.SHA256, data, base64sign)
		if err != nil {
			t.Fatal("Sha256WithRsaSignVerify fail:", err)
		}
	}
}
