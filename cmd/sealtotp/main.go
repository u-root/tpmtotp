package main

import (
	"log"

	"github.com/u-root/tpmtotp/pkg/token"
)

const ()

func main() {
	_, qrImage, err := token.CreateQRSecretHOTP()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	token.PrintTotpInAnsii(qrImage)

}
