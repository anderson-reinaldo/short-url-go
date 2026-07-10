package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"

	"github.com/anderson-reinaldo/short-url-go/src/config"
)

func Encrypt(originalURL string) (string, error) {
	block, err := aes.NewCipher([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(originalURL), nil)

	return hex.EncodeToString(cipherText), nil
}
