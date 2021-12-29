package api

import (
	"Exam/test3/service"
	"Exam/test3/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserName string `json:"userName"`
	jwt.StandardClaims
}

func Auth(c *gin.Context) {
	token, err1 := c.Cookie("token")
	if err1 != nil {
		fmt.Println("err1:", err1)
		tool.RespErrorWithDate(c, "没有token!")
		c.Abort()
		return
	}
	claim, err := service.ParseToken(token)

	num := tool.TokenCheck(claim, err)
	switch num {
	case 2:
		tool.RespErrorWithDate(c, "token错误！")
		c.Abort()
		return
	case 1:
		tool.RespErrorWithDate(c, "您的token过期啦！请重新登陆的啦~")
		c.Abort()
		return
	}

}

//Old
//Auth jwt鉴权
//func Auth(c *gin.Context) {
//	token, err1 := c.Cookie("token")
//	refreshToken, err2 := c.Cookie("refreshToken")
//	userName, err3 := c.Cookie("userName")
//	//因为获取cookie里面的token是设置的过期就消失，如果过期了可能检验不到，所以就先不检查err1
//	if err2 != nil || err3 != nil {
//		fmt.Println("err1:", err1, "err2:", err2)
//		tool.RespErrorWithDate(c, "error!")
//		c.Abort()
//		return
//	}
//	//如果token还没有过期，就先检验token
//	if err1 == nil {
//		//先解析token
//		claim, err := service.ParseToken(token)
//
//		//再检查token是否无效
//		num := tool.TokenCheck(c, claim, err)
//		fmt.Println(err)
//		if num == 2 {
//			tool.RespErrorWithDate(c, gin.H{
//				"error:": err,
//			})
//			c.Abort()
//			return
//		}
//	} else {
//
//		//涉及到过期问题，先检查refreshToken的情况
//		claim, err := service.ParseRefreshToken(refreshToken)
//
//		//再检查refreshToken下的claim是否错误
//		num := tool.TokenCheck(c, claim, err)
//		//老规矩，1是错误，2是过期，对于refreshToken过期或错误，都需要重新登录
//		if num == 1 || num == 2 {
//			tool.RespErrorWithDate(c, "您的登录已过期！请重新登陆")
//			c.Abort()
//			return
//		}
//		//正常就是0，就放行重新发放一个token
//		newToken, err4 := service.SetJWT(5, userName, "token")
//		fmt.Println(newToken)
//		c.SetCookie("token", newToken, 5, "/", "", false, true)
//		fmt.Println("已发放一个新的token啦~")
//		if err4 != nil {
//			tool.RespErrorWithDate(c, err4)
//			c.Abort()
//			return
//		}
//	}
//}
