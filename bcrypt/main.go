package main

import (
	"fmt"
)

func main() {

	var plainPwd = "cU1@sz5+yw-697~4onEiKd#MqpgLFN"
	fmt.Println(plainPwd)

	bcryptSecret, err := GenerateBcryptSecret(plainPwd)
	fmt.Println(bcryptSecret, err, len(bcryptSecret))

	ok := VerifyBcryptSecret(bcryptSecret, plainPwd)
	fmt.Println(ok)
}
