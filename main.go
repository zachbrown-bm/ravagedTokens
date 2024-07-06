package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"ravagedTokens/handlers"
)

func main() {
	serverPort := ":3010"

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("you found the root, congrats."))
	})

	router.Route("/token", handlers.TokenHandlers)

	fmt.Printf("Starting server on port: %s\n", serverPort)
	err := http.ListenAndServe(serverPort, router)
	if err != nil {
		fmt.Println(err)
	}
}
