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

func CreateToken(signMethod *jwt.SigningMethodHMAC, data jwt.Claims, hmacSecret []byte) (string, error) {

	token := jwt.NewWithClaims(signMethod, data)
	return token.SignedString(hmacSecret)

}

func ParseTokenWithClaims(tokenString string, hmacSecret []byte) (jwt.RegisteredClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return *claims, nil
	} else {
		return jwt.RegisteredClaims{}, err
	}

}

func ParseToken(tokenString string, hmacSecret []byte) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}
