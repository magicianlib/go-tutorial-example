package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {

	var secret = []byte("The secret key")
	var data = []byte("data")

	messageMac, err := HmacSHA512(secret, data)
	if err != nil {
		fmt.Println(err)
	}

	// Convert byte to hex string
	fmt.Println("hmacSha512:", hex.EncodeToString(messageMac))

	valid := HmacSHA512Valid(messageMac, secret, data)
	fmt.Println("Valid:", valid)
}

func HmacSHA512(secret, data []byte) (mac []byte, err error) {

	// hmacMd5: md5.New
	hash := hmac.New(sha512.New, secret)
	_, err = hash.Write(data)
	if err != nil {
		return nil, err
	}

	mac = hash.Sum(nil)
	// Convert byte to hex string: hex.EncodeToString(mac)

	return
}

func HmacSHA512Valid(expectedMac, secret, data []byte) bool {
	messageMac, err := HmacSHA512(secret, data)
	if err != nil {
		return false
	}

	return hmac.Equal(messageMac, expectedMac)
}
