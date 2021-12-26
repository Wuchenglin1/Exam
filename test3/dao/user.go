package dao

import (
	"Exam/test3/model"
)

// SelectUsername 返回一个user结构体和一个error
func SelectUsername(name string) (model.User, error) {
	user := model.User{}
	err := dB.QueryRow("select id,name,password from user where name = ?", name).Scan(&user.Id, &user.Username, &user.UserPassword)
	return user, err
}

func InsertUser(user model.User) error {
	_, err := dB.Exec("insert into user(name,password) values(?,?)", user.Username, user.UserPassword)
	return err
}
