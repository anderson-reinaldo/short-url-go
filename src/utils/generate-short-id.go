package utils

import (
	"crypto/rand"
	"log"
	"math/big"
)

var lettersRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateShortID() string {
	b := make([]rune, 6)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt((int64(len(lettersRunes)))))
		if err != nil {
			log.Fatal(err)
		}

		b[i] = lettersRunes[num.Uint64()]
	}

	return string(b)
}
