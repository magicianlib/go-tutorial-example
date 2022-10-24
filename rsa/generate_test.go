package rsa_example_test

import (
	"path/filepath"
	"testing"

	rsa_example "itknown.io/rsa"
)

func TestGenerateRsaKeyPair(t *testing.T) {
	abs, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	derDir := filepath.Join(abs, "der")

	t.Log(abs)

	priPEM, pubPEM, err := rsa_example.GenerateRsaKeyPair(4096, filepath.Join(derDir, "id.der"))
	if err != nil {
		panic(err)
	}

	t.Log(string(priPEM))
	t.Log(string(pubPEM))
}
