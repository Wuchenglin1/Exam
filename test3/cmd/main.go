package main

import (
	"Exam/test3/api"
	"Exam/test3/dao"
)

func main() {
	dao.InitMySql()
	api.InitEngine()
}
