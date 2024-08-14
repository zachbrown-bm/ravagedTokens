package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ravagedTokens/handlers"
)

func main() {
	serverPort := ":3010"

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/", handlers.RootHandler)

	tokens := e.Group("/token")
	tokens.GET("/", handlers.GetTokens)
	tokens.GET("/:tokenId", handlers.GetToken)
	tokens.DELETE("/:tokenId", handlers.DeleteToken)

	// user := e.Group("/user")
	// user.DELETE("/{userId}", handlers.DestroyUserHandler)
	e.DELETE("/user/:userId", handlers.DestroyUserHandler)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(serverPort))
}
