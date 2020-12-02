package sharedsecret

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// EncryptBytes can encrypt the given message, using a shared key (32 bytes long)
func EncryptBytes(msg []byte, key [32]byte) ([]byte, error) {
	// Ref. https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
	c, err := aes.NewCipher(key[:])
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return []byte{}, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}
	return gcm.Seal(nonce, nonce, msg, nil), nil
}

// Encrypt can encrypt the given message, using a shared password.
// The password string should ideally have a size of 32 bytes.
func Encrypt(message, password string) (string, error) {
	var key [32]byte
	copy(key[:], password)
	encryptedMessage, err := EncryptBytes([]byte(message), key)
	if err != nil {
		return "", err
	}
	return string(encryptedMessage), nil
}
