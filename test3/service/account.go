package service

import (
	"Exam/test3/dao"
	"Exam/test3/model"
)

func Account(userName string) model.User {
	iUser := dao.Account(userName)
	return iUser
}

func Transfer(t model.Transfer) (error, bool) {
	err, is := dao.Transfer(t)

	return err, is
}

func Commit() error {
	err := dao.Commit()
	return err
}

func RollBack() error {
	err := dao.RollBack()
	return err
}

func TransferSelect(k string) (map[int]model.Transfer, error) {
	m, err := dao.TransferSelect(k)
	return m, err
}

func TransferAddInfo(t model.Transfer) error {
	err := dao.TransferAddInfo(t)
	return err
}

func CZ(cz model.CZ) error {
	err := dao.CZ(cz)
	return err
}
