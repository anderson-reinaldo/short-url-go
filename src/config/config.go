package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	Port     = 0
	Secret   = ""
	BaseURL  = ""
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

	BaseURL = os.Getenv("BASE_URL")
	if BaseURL == "" {
		BaseURL = fmt.Sprintf("http://localhost:%d", Port)
	}
}

func MuLock() {
	mu.Lock()
}

func MuUnLock() {
	mu.Unlock()
}
