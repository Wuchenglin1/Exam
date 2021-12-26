package service

import (
	"Exam/test3/dao"
	"Exam/test3/model"
	"golang.org/x/crypto/bcrypt"
)

// IsRepeatUsername 如果账户不存在则返回一个err，若存在则返回一个nil
func IsRepeatUsername(username string) (model.User, error) {
	user, err := dao.SelectUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func Register(user model.User) error {
	//在此处将密码进行加盐哈希处理
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
	user.UserPassword = string(hash)
	err := dao.InsertUser(user)
	return err
}

// IsUserCorrect 检查账号密码是否正确，如果正确返回一个true，若错误则返回一个false
func IsUserCorrect(user *model.User) bool {
	iUser, _ := dao.SelectUsername(user.UserName)
	err := bcrypt.CompareHashAndPassword([]byte(iUser.UserPassword), []byte(user.UserPassword))
	if err == nil {
		user.Id = iUser.Id
		return true
	} else {
		return false
	}
}
