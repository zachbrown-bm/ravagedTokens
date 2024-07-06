package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"ravagedTokens/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("you found the root, congrats."))
	})

	router.Route("/token", handlers.TokenHandlers)

	err := http.ListenAndServe(":3010", router)
	if err != nil {
		fmt.Println(err)
	}
}
