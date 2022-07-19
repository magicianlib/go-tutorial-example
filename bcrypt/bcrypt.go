package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// go get -u golang.org/x/crypto/bcrypt

func GenerateBcryptSecret(plainPwd string) (string, error) {
	return GenerateBcryptSecretWithCost(plainPwd, bcrypt.DefaultCost)
}

// GenerateBcryptSecretWithCost Use bcrypt.GenerateFromPassword to hash & salt pwd.
//
// bcrypt.MinCost is just an integer constant provided by the bcrypt package
// along with bcrypt.DefaultCost & bcrypt.MaxCost.
//
// The cost can be any value you want provided it isn't lower
// than the MinCost(4) or more than the MinCost(31)
//
func GenerateBcryptSecretWithCost(plainPwd string, cost int) (string, error) {

	if cost > bcrypt.MaxCost || cost < bcrypt.MinCost {
		cost = bcrypt.DefaultCost
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(plainPwd), cost)
	if err != nil {
		log.Printf("bcrypt generate password [%s] err: %v\n", plainPwd, err)
		return "", err
	}

	// GenerateFromPassword returns a byte slice, so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

func VerifyBcryptSecret(bcryptSecret string, plainPwd string) bool {

	// Since we'll be getting the hashed password from the DB it
	// will be a string. So we'll need to convert it to a byte slice

	err := bcrypt.CompareHashAndPassword([]byte(bcryptSecret), []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
