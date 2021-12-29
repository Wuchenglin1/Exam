package service

import (
	"Exam/test3/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// ParseToken 解析普通的Token
func ParseToken(tokenStr string) (*model.Claims, error) {
	JWT := model.JWT{}
	signedStr := []byte(JWT.SigningKey)
	token, err := jwt.ParseWithClaims(tokenStr, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return signedStr, nil
	})

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// ParseRefreshToken 解析refreshToken
//func ParseRefreshToken(tokenStr string) (*model.Claims, error) {
//	JWT := model.JWT{}
//	signedStr := []byte(JWT.SigningKey)
//	token, err := jwt.ParseWithClaims(tokenStr, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return signedStr, nil
//	})
//
//	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
//		if claims.TokenType == "token" {
//			errClaims := new(model.Claims)
//			errClaims.TokenType = "err"
//			return errClaims, nil
//		}
//		return claims, nil
//	} else {
//		return nil, err
//	}
//}

//SetJWT 返回一个有效时间为ExpireTime的JWT和error
func SetJWT(ExpireTime int64, userName string, tokenType string) (string, error) {
	//将user信息和tokenType先写进Claim里
	//因为payload是没有经过加密处理的，任何人都可以解密里面的信息，所以里面不存放私密信息
	JWT := model.JWT{}
	signedStr := []byte(JWT.SigningKey)
	//因为payload可以被解密查看，所以在payload里面只存放了一个userName+当前时间戳的字符串
	str := userName + strconv.FormatInt(time.Now().Unix(), 10)

	//生成一个Claim
	myClaim := model.Claims{
		UserStr: str,
		//TokenType: tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + ExpireTime,
		},
	}

	//再将Claim用HS256加密生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

	//最后签名得到jwt
	s, err := token.SignedString(signedStr)
	if err != nil {
		fmt.Println(err)
		return s, err
	}
	return s, nil
}
