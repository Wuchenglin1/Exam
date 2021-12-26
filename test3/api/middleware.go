package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type Claims struct {
	UserName string `json:"userName"`
	jwt.StandardClaims
}

var secretKey = []byte("SecretKey") //模拟私钥

//Auth jwt鉴权
func Auth(c *gin.Context) {
	userName, _ := c.Cookie("userName")
	//完善payload信息，将userName信息填充进去
	claim := Claims{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "wuchenglin",
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
		},
	}
	//登录之后获取存储在cookie中的jwt信息
	t, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(403, "您还没有登录！")
		c.Abort()
		return
	}
	//对比jwt
	token, err1 := jwt.ParseWithClaims(t, &claim, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	//token.Valid(bool)检验签名是否有效
	if token.Valid {
		c.JSON(200, gin.H{
			"您好！": claim.UserName,
		})
	} else if ve, ok := err1.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			//ValidationErrorMalformed验证token是否为畸形
			//token的格式错误,
			c.JSON(403, "请输个像样的token")
			c.Abort()
			return
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			//jwt.ValidationErrorExpired是验证jwt签名是否过期
			//jwt.ValidationErrorNotValidYet是验证用户操作是否活跃
			c.JSON(403, "您不活跃或者验证已过期！")
			c.Abort()
			return
		} else {
			c.JSON(403, gin.H{
				"error:": err1,
			})
		}
	} else {
		c.JSON(403, gin.H{
			"不能识别此token": err1,
		})
	}
}
