package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
	"todo-api/config"
	"todo-api/model"
)

func GenerateJwtToken(user *model.User) (string, error) {
	expTimeMs, _ := strconv.Atoi(config.New().JWTExpirationMs)
	exp := time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix()
	name := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	id := strconv.Itoa(user.ID)

	// Set custom claims
	claims := &model.JwtCustomClaims{
		id,
		name,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	jwt, err := token.SignedString([]byte(config.New().JWTSecret))
	return jwt, err
}

func GetUserIdFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	return claims.ID
}
