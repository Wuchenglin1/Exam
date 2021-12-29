package api

import (
	"Exam/test3/dao"
	"github.com/gin-gonic/gin"
	"time"
)

func InitEngine() {
	engine := gin.Default()

	//账户自动扣费的接口,每5分钟扣一块钱
	go func() {
		for {
			time.Sleep(time.Minute * 5)
			dao.Deduction()
		}
	}()

	user := engine.Group("/user")
	{
		user.POST("/register", register)
		user.POST("/login", Login)
	}

	account := engine.Group("/account", Auth)
	{
		account.GET("", Account)
		account.PUT("/cz", CZ)
		account.POST("/transfer", Transfer)
		account.POST("/transferSelect", TransferSelect)
		account.PUT("/transferAddInfo", TransferAddInfo)
	}

	_ = engine.Run()
}
