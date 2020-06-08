package main

import (
	"log"

	"github.com/u-root/tpmtotp/pkg/token"
	"github.com/u-root/u-root/pkg/tss"
)

const ()

func main() {
	tpm, err := tss.NewTPM()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	_, _, err = token.CreateQRSecretHOTP()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	err = tpm.TakeOwnership("", "")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
