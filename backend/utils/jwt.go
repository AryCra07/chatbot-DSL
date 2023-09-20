package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	JwtKeyString    string        = "AryCra07-IfWinterComesCanSpringBeFarBehind?"
	TokenValidHours time.Duration = 12
)

var JwtKey = []byte(JwtKeyString)

type Claims struct {
	UserId int
	Auth   int
	jwt.StandardClaims
}

func ReleaseToken(userId int, auth int) (string, error) {
	// token expire time
	expirationTime := time.Now().Add(TokenValidHours * time.Hour)

	claims := &Claims{
		// self design options
		UserId: userId,
		Auth:   auth,
		// basic options
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "AryCra07",
			Subject:   "user token",
		},
	}

	// jwt generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	return token, claims, err
}
