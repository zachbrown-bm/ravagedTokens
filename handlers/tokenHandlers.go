package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func GetTokens(c echo.Context) error {
	_, err := c.Response().Write([]byte(fmt.Sprint("Here's a list of tokens")))
	if err != nil {
		return err
	}

	return nil
}

func GetToken(c echo.Context) error {
	tokenId := c.Param("tokenId")
	_, err := c.Response().Write([]byte(fmt.Sprintf("Here's token: %s", tokenId)))
	if err != nil {
		return err
	}

	return nil
}

func DeleteToken(c echo.Context) error {
	tokenId := c.Param("tokenId")
	_, err := c.Response().Write([]byte(fmt.Sprintf("Removed Token: %s", tokenId)))
	if err != nil {
		return err
	}

	return nil
}
