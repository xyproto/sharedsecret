package main

import (
	"fmt"
	"os"

	"github.com/xyproto/ask"
	"github.com/xyproto/sharedsecret"
)

const password = "hunter1"

func main() {
	message := ask.Ask("Type in a short message: ")
	encryptedMessage, err := sharedsecret.Encrypt(message, password)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("The encrypted message is: % x\n", encryptedMessage)
	decryptedMessage, err := sharedsecret.Decrypt(encryptedMessage, password)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("The decrypted message is: %s\n", decryptedMessage)
}
