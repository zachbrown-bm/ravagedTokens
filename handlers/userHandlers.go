package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"log"
	"ravagedTokens/framework/usermanager"
	"strconv"
)

func DestroyUserHandler(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		uid = -1
	}

	currentUser := usermanager.GetCurrent()
	currentUser.Id = uid - 1
	currentUser.Email = "anew@email.com"

	// loading and parsing templates, do this once for better performance
	// then use parsed templates multiple times
	tmpl, err := template.New("").ParseGlob("html/**/*.gohtml")
	if err != nil {
		log.Println("Error parsing templates", err)
		return err
	}

	// executing template named "homepage"
	if err := tmpl.ExecuteTemplate(c.Response().Writer, "user", currentUser); err != nil {
		fmt.Println("Error executing template for root handler.", err)
		return err
	}

	return nil
}
