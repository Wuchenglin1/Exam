package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserStr string
	//TokenType string //token,refreshToken,err
	Time time.Time
	jwt.StandardClaims
}

type JWT struct {
	SigningKey string `json:"signingKey"`
}
