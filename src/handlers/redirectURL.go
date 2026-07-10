package handlers

import (
	"net/http"

	"github.com/anderson-reinaldo/short-url-go/src/config"
	"github.com/anderson-reinaldo/short-url-go/src/utils"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortID := r.URL.Path[1:]

	config.MuLock()
	encryptedURL, ok := config.UrlStore[shortID]
	config.MuUnLock()

	if !ok {
		http.Error(w, "A URL que você está buscando não existe.", http.StatusNotFound)
		return
	}

	decryptedURL, err := utils.Decrypt(encryptedURL)
	if err != nil {
		http.Error(w, "Não foi possível processar esta URL.", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, decryptedURL, http.StatusFound)

}
