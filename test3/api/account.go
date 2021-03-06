package api

import (
	"Exam/test3/model"
	"Exam/test3/service"
	"Exam/test3/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"unicode"
)

//Account 查看自己账户有多少钱(若第一次使用则自动开户)
func Account(c *gin.Context) {
	user := model.User{}
	//因为用户已经登录，这里就不处理error了
	user.Id, _ = c.Cookie("id")
	user.UserName, _ = c.Cookie("userName")
	iUser := service.Account(user.UserName)
	tool.RespSuccessfullWithDate(c, gin.H{
		"time:":                time.Now(),
		"尊敬的" + iUser.UserName: "您好！",
		"您的余额为：":               iUser.Money,
	})
}

// Transfer 需要三个关键词toWhom(给谁转账),detail(转账详情),money(转账的金额)
func Transfer(c *gin.Context) {
	mUser := model.User{}
	//因为用户已经登录，这里就不处理error了
	mUser.Id, _ = c.Cookie("id")
	mUser.UserName, _ = c.Cookie("userName")
	//查询登录用户拥有多少钱
	mUser = service.Account(mUser.UserName)
	fmt.Println(mUser)
	user := model.User{}
	//因为用户已经登录，这里就不处理error了
	user.UserName, _ = c.Cookie("userName")
	t := model.Transfer{
		UserName: user.UserName,
		ToWhom:   c.PostForm("toWhom"),
		Detail:   c.PostForm("detail"),
	}
	_, err1 := service.IsRepeatUsername(t.ToWhom)
	if err1 != nil {
		tool.RespErrorWithDate(c, "您想发送的用户不存在或还没有开户！")
	}

	if t.UserName == t.ToWhom {
		tool.RespErrorWithDate(c, "拜托！别开这种玩笑刷钱好不好得啦~")
		return
	}
	m, err := strconv.Atoi(c.PostForm("money"))
	if err != nil || c.PostForm("money") == "" {
		tool.RespErrorWithDate(c, "请正确输入转账的数字")
		return
	}
	t.Money = m
	if mUser.Money < t.Money {
		tool.RespErrorWithDate(c, "您的余额不足！")
		return
	}
	if t.Money < 0 {
		tool.RespErrorWithDate(c, "money 不能为负数！")
		return
	}

	err2, is := service.Transfer(t)
	if err2 != nil {
		tool.RespErrorWithDate(c, err)
		return
	}
	if is {
		c.JSON(200, "您是否确定转账？【请添加一个sure,0为否，1为确定】")
	}
	sure := c.PostForm("sure")
	if sure == "1" {
		err = service.Commit()
		if err != nil {
			fmt.Println(err)
			return
		}
		tool.RespSuccessfullWithDate(c, "恭喜您转账成功！")
	} else {
		if sure == "0" {
			err = service.RollBack()
			if err != nil {
				fmt.Println(err)
				return
			}
			_ = service.RollBack()
			tool.RespErrorWithDate(c, "转账失败了~")
		}
	}
}

//TransferSelect 需要添加一个关键词key(模糊查询的关键词)
func TransferSelect(c *gin.Context) {
	key := c.PostForm("key")
	m, err := service.TransferSelect(key)
	if err != nil {
		tool.RespErrorWithDate(c, "没有该记录！")
	}
	fmt.Println(m)
	for _, v := range m {
		tool.RespSuccessfullWithDate(c, v)
	}
}

func TransferAddInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithDate(c, "输入出错！")
		return
	}
	name, err2 := c.Cookie("userName")
	if err2 != nil {
		tool.RespErrorWithDate(c, "没有cookie！")
		return
	}
	t := model.Transfer{
		Id:       id,
		UserName: name,
		Detail:   c.PostForm("key"),
	}
	n := func(str string) (count int) {
		for _, char := range str {
			if unicode.Is(unicode.Han, char) {
				count++
			}
		}
		return
	}(t.Detail)
	if n >= 20 {
		tool.RespErrorWithDate(c, "您输入的描述太长啦！请重新输入的啦~")
		return
	}
	err = service.TransferAddInfo(t)
	if err != nil {
		tool.RespErrorWithDate(c, "添加失败！")
	}
	tool.RespSuccessfullWithDate(c, "修改成功！")
}

func CZ(c *gin.Context) {
	userName, err := c.Cookie("userName")
	if err != nil {
		tool.RespErrorWithDate(c, "没有cookie!")
		c.Abort()
		return
	}
	m, err1 := strconv.Atoi(c.PostForm("money"))
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	cz := model.CZ{
		UserName: userName,
		Money:    m,
	}

	err = service.CZ(cz)
	if err != nil {
		fmt.Println(err)
		tool.RespErrorWithDate(c, "充值失败！")
		return
	}
	tool.RespSuccessfullWithDate(c, cz.UserName+"充值"+strconv.Itoa(cz.Money)+"成功！")
}
