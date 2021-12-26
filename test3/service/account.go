package service

import (
	"Exam/test3/dao"
	"Exam/test3/model"
)

func Account(userName string) model.User {
	iUser := dao.Account(userName)
	return iUser
}

func Transfer(t model.Transfer) error {
	err := dao.Transfer(t)
	return err
}
