package handlers

import (
	"html/template"
	"net/http"
	"time"
)

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	data := map[string]interface{}{
		"PageTitle": "Welcome to Ravaged Tokens",
		"Date":      time.Now().Local().Format("2006-Jan-02"),
	}

	// y, m, d = now.Date()		// can use this to get different parts if I want to format differently

	tmpl, err := template.ParseFiles("html/masters/index.html")
	if err != nil {
		w.Write([]byte("Couldn't find template file for root handler."))
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		w.Write([]byte("Error parsing template for root handler."))
	}
}
