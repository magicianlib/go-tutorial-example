package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// go get -u github.com/golang-jwt/jwt/v4

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
