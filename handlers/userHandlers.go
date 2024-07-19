package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ravagedTokens/framework/usermanager"
)

func DestroyUserHandler(w http.ResponseWriter, _ *http.Request) {
	currentUser := usermanager.GetCurrent()
	currentUser.Email = "anew@email.com"

	// loading and parsing templates, do this once for better performance
	// then use parsed templates multiple times
	tmpl, err := template.New("").ParseGlob("html/**/*.gohtml")
	if err != nil {
		log.Println("Error parsing templates", err)
		return
	}

	// executing template named "homepage"
	if err := tmpl.ExecuteTemplate(w, "user", currentUser); err != nil {
		fmt.Println("Error executing template for root handler.", err)
		return
	}
}
