package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/anderson-reinaldo/short-url-go/src/config"
	"github.com/anderson-reinaldo/short-url-go/src/utils"
)

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

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

	encryptURL, err := utils.Encrypt(originalURL)
	if err != nil {
		http.Error(w, "Não foi possível encurtar a URL.", http.StatusInternalServerError)
		return
	}

	config.MuLock()
	var shortID string
	for {
		shortID = utils.GenerateShortID()
		if _, exists := config.UrlStore[shortID]; !exists {
			break
		}
	}
	config.UrlStore[shortID] = encryptURL
	config.MuUnLock()

	shortURL := fmt.Sprintf("%s/%s", config.BaseURL, shortID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{ShortURL: shortURL})
}
