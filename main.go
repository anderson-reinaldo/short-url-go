package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anderson-reinaldo/short-url-go/src/config"
	"github.com/anderson-reinaldo/short-url-go/src/handlers"
)

func init() {
	config.Load()
}

func main() {
	http.HandleFunc("/shorten", handlers.ShortenURLHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	fmt.Printf("Server is running on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
