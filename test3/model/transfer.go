package model

type Transfer struct {
	Id       int
	UserName string
	ToWhom   string
	Money    int
	Detail   string
}

type CZ struct {
	UserName string
	Money    int
}
