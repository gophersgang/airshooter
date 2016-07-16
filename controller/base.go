package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"net/http"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("JWT Auth")
		user := c.Get("user").(*jwt.Token)
		// DBに登録されたtokenと照合して、validなclientかどうか検証する
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		if name != "Jon Snow" {
			return c.String(http.StatusOK, "Failed "+name+"")
		}
		return next(c)
	}
}

func ValidationHandler(target interface{}) error {
	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)
	if err := validate.Struct(target); err != nil {
		return err
	}

	return nil
}
