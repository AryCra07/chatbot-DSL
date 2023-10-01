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
	UserId   string `json:"userId"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(userId string, password string) (string, error) {
	// token expire time
	expirationTime := time.Now().Add(TokenValidHours * time.Hour)

	claims := &Claims{
		// self design options
		UserId:   userId,
		Password: password,
		// basic options
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "AryCra07",
			Subject:   "user token",
		},
	}

	// jwt generate token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtKey)

	return token, err
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	return token, claims, err
}
