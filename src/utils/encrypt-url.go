package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/anderson-reinaldo/short-url-go/src/config"
)

func Encrypt(originalURL string) string {
	block, err := aes.NewCipher([]byte(config.Secret))
	if err != nil {
		log.Fatal(err)
	}

	plainText := []byte(originalURL)
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]

	if _, err := rand.Read(iv); err != nil {
		log.Fatal(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(cipherText)
}
