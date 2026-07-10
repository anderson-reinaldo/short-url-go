package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anderson-reinaldo/short-url-go/src/config"
	"github.com/anderson-reinaldo/short-url-go/src/utils"
)

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		http.Error(w, "Parametro URL no query é obrigatorio!", http.StatusBadRequest)
		return
	}

	if !(strings.HasPrefix(originalURL, "https://") || strings.HasPrefix(originalURL, "http://")) {
		http.Error(w, "Parametro URL no query precisa ter o protocolo https:// ou http:// antecedendo o valor da URL", http.StatusBadRequest)
		return
	}

	encryptURL := utils.Encrypt(originalURL)
	shortID := utils.GenerateShortID()
	config.MuLock()
	config.UrlStore[shortID] = encryptURL
	defer config.MuUnLock()

	shortURL := fmt.Sprintf("http://localhost:5000/%s", shortID)
	fmt.Fprintf(w, "A URL encurtada desta url original é: %s", shortURL)

}
