package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器错误",
	})
}

func RespErrorWithDate(c *gin.Context, date interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"info": date,
	})
}

//func RespSuccessful(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"info": "成功",
//	})
//}

func RespSuccessfullWithDate(c *gin.Context, date interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"date": date,
	})
}
