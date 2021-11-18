package middlewares

import (
	"kemejaku/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func GenerateTokenJWT(userId int) (string, error) {

	// Set custom claims atau payload
	claims := &JwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims atau header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(constants.SECRET_JWT))
	return t, err
}
