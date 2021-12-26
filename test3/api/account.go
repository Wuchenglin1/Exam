package api

import (
	"Exam/test3/model"
	"Exam/test3/service"
	"Exam/test3/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

//Account 查看自己账户有多少钱(若第一次使用则自动开户)
func Account(c *gin.Context) {
	user := model.User{}
	//因为用户已经登录，这里就不处理error了
	user.Id, _ = c.Cookie("id")
	user.UserName, _ = c.Cookie("userName")
	iUser := service.Account(user.UserName)
	tool.RespSuccessfullWithDate(c, gin.H{
		"尊敬的" + iUser.UserName: "您好！",
		"您的余额为：":               iUser.Money,
	})
}

func Transfer(c *gin.Context) {
	mUser := model.User{}
	//因为用户已经登录，这里就不处理error了
	mUser.Id, _ = c.Cookie("id")
	mUser.UserName, _ = c.Cookie("userName")
	yUser := service.Account(mUser.UserName)

	user := model.User{}
	//因为用户已经登录，这里就不处理error了
	user.UserName, _ = c.Cookie("userName")
	t := model.Transfer{
		UserName: user.UserName,
		ToWhom:   c.PostForm("toWhom"),
		Detail:   c.PostForm("detail"),
	}
	m, err := strconv.Atoi(c.PostForm("money"))
	if err != nil {
		tool.RespErrorWithDate(c, "请正确输入转账的数字")
		return
	}
	if yUser.Money < t.Money {
		tool.RespErrorWithDate(c, "您的余额不足！")
		return
	}
	t.Money = m

	err1 := service.Transfer(t)
	if err1 != nil {
		tool.RespErrorWithDate(c, err)
		return
	}
}
