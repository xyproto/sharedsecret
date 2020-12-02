package sharedsecret

import (
	"crypto/aes"
	"crypto/cipher"
)

func DecryptBytes(data []byte, key [32]byte) ([]byte, error) {
	c, err := aes.NewCipher(key[:])
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return []byte{}, err
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return []byte{}, err
	}
	nonce, data := data[:nonceSize], data[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return []byte{}, err
	}
	return decrypted, nil
}

func Decrypt(message, password string) (string, error) {
	var key [32]byte
	copy(key[:], password)
	decryptedMessage, err := DecryptBytes([]byte(message), key)
	if err != nil {
		return "", err
	}
	return string(decryptedMessage), nil
}
