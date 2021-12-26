package api

import (
	"Exam/test3/model"
	"Exam/test3/service"
	"Exam/test3/tool"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

//需要传入两个key:username和password
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) <= 6 || len(password) <= 6 {
		tool.RespErrorWithDate(c, "非法输入：账号和密码长度需要大于六位！")
		return
	}
	user := model.User{
		UserName:     username,
		UserPassword: password,
	}

	err := service.Register(user)
	if err != nil {
		fmt.Println("ERROR1", err)
		tool.RespErrorWithDate(c, "注册失败，账号已存在！")
		return
	}
	tool.RespSuccessfullWithDate(c, "注册成功！")
}

// Login 需要传入两个key:username和password
func Login(c *gin.Context) {
	user := model.User{
		UserName:     c.PostForm("username"),
		UserPassword: c.PostForm("password"),
	}
	_, err := service.IsRepeatUsername(user.UserName)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("ERROR:", err)
			tool.RespInternalError(c)
			return
		} else {
			tool.RespErrorWithDate(c, "账号不存在！")
			return
		}
	}

	is := service.IsUserCorrect(&user)

	if is == false {
		tool.RespErrorWithDate(c, "您的密码错误!")
		return
	}
	fmt.Println(user)
	j, err1 := tool.SetJWT(c, user.UserName)
	if err1 != nil {
		tool.RespErrorWithDate(c, err1)
		return
	}
	c.SetCookie("userName", user.UserName, 3600, "/", "", false, true)
	c.SetCookie("id", user.Id, 3600, "/", "", false, true)
	c.SetCookie("jwt", j, 3600, "/", "", false, true)

	tool.RespSuccessfullWithDate(c, "登录成功！")

}
