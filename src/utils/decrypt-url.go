package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"

	"github.com/anderson-reinaldo/short-url-go/src/config"
)

func Decrypt(encryptedURL string) string {
	block, err := aes.NewCipher([]byte(config.Secret))
	if err != nil {
		log.Fatal(err)
	}

	cipherText, err := hex.DecodeString(encryptedURL)
	if err != nil {
		log.Fatal(err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)

}
