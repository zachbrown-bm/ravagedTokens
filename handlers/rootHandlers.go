package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ravagedTokens/framework/usermanager"
	"time"
)

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	currentUser := usermanager.GetCurrent()
	data := map[string]interface{}{
		"PageTitle": "Welcome to Ravaged Tokens",
		"Date":      time.Now().Local().Format("2006-Jan-02"),
		"User":      currentUser,
	}

	// loading and parsing templates, do this once for better performance
	// then use parsed templates multiple times
	tmpl, err := template.New("").ParseGlob("html/**/*.gohtml")
	if err != nil {
		log.Println("Error parsing templates", err)
		return
	}

	// executing template named "homepage"
	if err := tmpl.ExecuteTemplate(w, "index", data); err != nil {
		fmt.Println("Error executing template for root handler.", err)
		return
	}
}
