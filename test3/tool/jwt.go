package tool

import (
	"Exam/test3/model"
)

// TokenCheck 这里只是简单地判断下这个claims是否正确。0：无错，1：过期，2：错误
func TokenCheck(claims *model.Claims, err error) int {
	if err != nil {
		if err.Error()[:16] == "token is expired" {
			return 1
		}
		return 2
	}
	return 0
}

//Old
//这里是想用refreshToken的，但是单独用cookie来实现好麻烦啊！想不出来了，就等以后合作的时候再用单独的路由来做吧！QAQ
////TokenCheck 检查Token是否错误或过期,有错误返回2,过期返回1,无错误返回0
//func TokenCheck(c *gin.Context, claims *model.Claims, err error) int {
//	//在解析的时候如果token错了TokenType会是err
//	if err == nil && claims.TokenType == "err" {
//		return 2
//	}
//	if err != nil {
//		//这里借鉴的达达子的鉴别token是否过期的方法,检测错误提示的前16个字符
//		if err.Error()[:16] == "token is expired" {
//			return 1
//		}
//		fmt.Println(err)
//		return 2
//	}
//
//	return 0
//}
