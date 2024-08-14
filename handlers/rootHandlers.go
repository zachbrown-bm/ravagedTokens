package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"log"
	"ravagedTokens/framework/usermanager"
	"time"
)

func RootHandler(c echo.Context) error {
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
		return err
	}

	// executing template named "homepage"
	if err := tmpl.ExecuteTemplate(c.Response().Writer, "index", data); err != nil {
		fmt.Println("Error executing template for root handler.", err)
		return err
	}

	return nil
}
