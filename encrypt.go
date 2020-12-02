package sharedsecret

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func EncryptBytes(msg []byte, key [32]byte) ([]byte, error) {
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

func Encrypt(message, password string) (string, error) {
	var key [32]byte
	copy(key[:], password)
	encryptedMessage, err := EncryptBytes([]byte(message), key)
	if err != nil {
		return "", err
	}
	return string(encryptedMessage), nil
}
