package dao

import (
	"Exam/test3/model"
	"fmt"
)

func Account(userName string) model.User {
	iUser := model.User{
		UserName: userName,
	}
	err := dB.QueryRow("select id,money from account where name = ?", iUser.UserName).Scan(&iUser.Id, &iUser.Money)
	if err != nil {
		err = dB.QueryRow("select id from user where name = ?", iUser.UserName).Scan(&iUser.Id)
		if err != nil {
			fmt.Println(err)
		}
		_, err = dB.Exec("insert into account(id, name, money) values(?,?,?) ", iUser.Id, iUser.UserName, 0)
		if err != nil {
			fmt.Println(err)
		}
	}
	return iUser
}

func Transfer(t model.Transfer) error {
	_, err := dB.Exec("begin ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = dB.Exec("insert into transfer (whom, toWhom, money, details) values (?,?,?,?)", t.UserName, t.ToWhom, t.Money, t.Detail)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
