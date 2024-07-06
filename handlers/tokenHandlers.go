package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func TokenHandlers(r chi.Router) {
	r.Get("/", getTokens)

	r.Get("/{tokenId}", getToken)
	r.Delete("/{tokenId}", deleteToken)
}

func getTokens(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(fmt.Sprint("Here's a list of tokens")))
}

func getToken(w http.ResponseWriter, r *http.Request) {
	tokenId := chi.URLParam(r, "tokenId")
	w.Write([]byte(fmt.Sprintf("Here's token: %s'", tokenId)))
}

func deleteToken(w http.ResponseWriter, r *http.Request) {
	tokenId := chi.URLParam(r, "tokenId")
	w.Write([]byte(fmt.Sprintf("Removed Token: %s", tokenId)))
}
