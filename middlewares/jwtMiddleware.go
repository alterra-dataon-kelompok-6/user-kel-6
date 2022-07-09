package middlewares

import (
	"errors"
	"log"
	"strings"
	"time"
	"user-kel-6/constants"

	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func GetUserIdFromToken(e echo.Context) uint {
	authToken := e.Request().Header.Get("Authorization")
	log.Println("01 - authToken", authToken)
	if authToken == "" {
		return 0
	}

	tokenString := strings.Split(authToken, " ")[1]
	log.Println("02 - tokenString", tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error token")
		}
		return []byte(constants.SECRET_JWT), nil
	})

	log.Println("03 - token", token)
	if !token.Valid || err != nil {
		log.Println("tidak valid ternyata :)")
		return 0
	}

	userId := token.Claims.(jwt.MapClaims)["userId"].(float64)
	log.Println("04", userId)
	if userId != 0 {
		return uint(userId)
	}
	return 0
}
