package tool

import (
	"Exam/test3/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var secretKey = []byte("SecretKey") //模拟私钥

func SetJWT(c *gin.Context, userName string) (string, error) {
	myClaim := model.Claims{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "wuchenglin",
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

	s, err2 := token.SignedString(secretKey)
	if err2 != nil {
		RespErrorWithDate(c, err2)
		return s, err2
	}
	return s, nil
}
