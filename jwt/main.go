package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var hmacSecret = []byte("nT2^&5#q")

func main() {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	claims := jwt.RegisteredClaims{
		Issuer:    "MG",
		Subject:   "",
		Audience:  nil,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 10)), // ExpiresAt 10s
		NotBefore: nil,
		IssuedAt:  nil,
		ID:        "94B66881-77A8-479F-9208-AE7D71DEEB60",
	}

	token, err := CreateToken(jwt.SigningMethodHS512, claims, hmacSecret)
	fmt.Printf("%#v\n%v\n\n", token, err)

	parseClaims, err := ParseTokenWithClaims(token, hmacSecret)
	fmt.Printf("%#v\n%v\n\n", parseClaims, err)

	time.Sleep(15 * time.Second)

	parseClaims, err = ParseTokenWithClaims(token, hmacSecret)
	fmt.Printf("%#v\n%v\n\n", parseClaims, err)
}
