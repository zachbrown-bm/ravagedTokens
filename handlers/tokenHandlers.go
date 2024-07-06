package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func TokenHandlers(r chi.Router) {
	r.Get("/", getToken)
	r.Delete("/{tokenId}", deleteToken)
}

func getToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprint("Here's a Token")))
}

func deleteToken(w http.ResponseWriter, r *http.Request) {
	tokenId := chi.URLParam(r, "tokenId")
	w.Write([]byte(fmt.Sprintf("Removed Token: %s", tokenId)))
}
