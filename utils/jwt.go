package utils

import (
	"net/http"
	"time"

	"github.com/Yahya-Elamri/signeitfaster/module"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(module.JwtSecret)
}

func Authorize(c echo.Context) (uint, bool) {
	cookie, err := c.Cookie("Token")
	if err != nil {
		if err == http.ErrNoCookie {
			return 0, false
		}
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return module.JwtSecret, nil
	})

	if err != nil {
		return 0, false
	}

	if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
		return claims.UserId, true
	}
	return 0, false
}
