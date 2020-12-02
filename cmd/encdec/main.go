package main

import (
	"fmt"
	"log"

	"github.com/xyproto/ask"
	"github.com/xyproto/sharedsecret"
)

func main() {
	const password = "hunter1"
	message := ask.Ask("Type in a short message: ")
	encryptedMessage, err := sharedsecret.Encrypt(message, password)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("The encrypted message is: % x\n", encryptedMessage)
	decryptedMessage, err := sharedsecret.Decrypt(encryptedMessage, password)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("The decrypted message is: %s\n", decryptedMessage)
}
