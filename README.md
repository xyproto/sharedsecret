# sharedsecret

Have a shared secret. Send encrypted messages.

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/xyproto/ask"
	"github.com/xyproto/sharedsecret"
)

const masterPassword = "hunter1"

func main() {
	message := ask.Ask("Type in a short message: ")
	encryptedMessage, err := sharedsecret.Encrypt(message, masterPassword)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("The encrypted message is: % x\n", encryptedMessage)
	decryptedMessage, err := sharedsecret.Decrypt(encryptedMessage, masterPassword)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("The decrypted message is: %s\n", decryptedMessage)
}
```

Interaction with the example program where the typed in message is "hi".

```
Type in a short message: hi
The encrypted message is: bf 18 25 16 d8 c1 de f0 70 4e 02 bd ec 20 6c 8b d6 8c f3 00 65 b5 fd d9 d6 06 f1 f9 ac 26
The decrypted message is: hi
```

## General info

* Version: 1.0.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
