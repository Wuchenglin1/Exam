package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()
	user := engine.Group("/user")
	{
		user.POST("/register", register)
		user.POST("/login", Login)
	}

	account := engine.Group("/account", Auth)
	{
		account.GET("", Account)
		account.POST("/transfer", Transfer)
		account.POST("/transferSelect", TransferSelect)
	}

	_ = engine.Run()
}
