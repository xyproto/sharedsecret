package sharedsecret

import (
	"crypto/aes"
	"crypto/cipher"
)

// DecryptBytes an decrypt the given message, if the correct shared key (32 bytes long) is given
func DecryptBytes(msg []byte, key [32]byte) ([]byte, error) {
	// Ref. https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
	c, err := aes.NewCipher(key[:])
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return []byte{}, err
	}
	nonceSize := gcm.NonceSize()
	if len(msg) < nonceSize {
		return []byte{}, err
	}
	nonce, data := msg[:nonceSize], msg[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return []byte{}, err
	}
	return decrypted, nil
}

// Decrypt can decrypt the given message, if the correct shared password is given
func Decrypt(message, password string) (string, error) {
	var key [32]byte
	copy(key[:], password)
	decryptedMessage, err := DecryptBytes([]byte(message), key)
	if err != nil {
		return "", err
	}
	return string(decryptedMessage), nil
}
