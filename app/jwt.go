package app

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type JWTClaim struct {
	ID       int
	Username string
	jwt.StandardClaims
}

func JsonWebToken(id int, username string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	SECRET := os.Getenv("JWT_SECRET")

	var JWT_KEY = []byte(SECRET)

	expTime := time.Now().Add(time.Hour * 24)
	claims := &JWTClaim{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "MyAuthService",
			ExpiresAt: expTime.Unix(),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}
	return token, nil
}
