package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	Port     = 0
	Secret   = ""
	UrlStore = make(map[string]string)
	mu       sync.Mutex
)

func Load() {
	var error error
	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("PORT"))
	if error != nil {
		Port = 5000
	}

	Secret = os.Getenv("SECRET")

}

func MuLock() {
	mu.Lock()
}

func MuUnLock() {
	mu.Unlock()
}
