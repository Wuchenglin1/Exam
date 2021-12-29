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

var m1, m2 int

func Transfer(t model.Transfer) (error, bool) {
	_, err := dB.Exec("begin ")
	if err != nil {
		fmt.Println(err)
		return err, false
	}

	_, err = dB.Exec("insert into transfer (whom, toWhom, money, details) values (?,?,?,?)", t.UserName, t.ToWhom, t.Money, t.Detail)
	if err != nil {
		fmt.Println(err)
		return err, false
	}

	err1 := dB.QueryRow("select money from account where name = ?", t.UserName).Scan(&m1)
	if err1 != nil {
		fmt.Println(err)
		return err, false
	}
	err1 = dB.QueryRow("select money from account where name = ?", t.ToWhom).Scan(&m2)
	if err1 != nil {
		fmt.Println(err)
		return err, false
	}
	m1 = m1 - t.Money
	m2 = m2 + t.Money

	_, err = dB.Exec("update account set money = ? where name = ?", m1, t.UserName)
	if err != nil {
		fmt.Println(err)
		return err, false
	}
	_, err = dB.Exec("update account set money = ? where name = ?", m2, t.ToWhom)
	if err != nil {
		fmt.Println(err)
		return err, false
	}

	return nil, true
}

func Commit() error {
	_, err := dB.Exec("commit ")
	return err
}

func RollBack() error {
	_, err := dB.Exec("rollback ")
	return err
}

func TransferSelect(k string) (map[int]model.Transfer, error) {
	i := 0
	m := make(map[int]model.Transfer)
	d := model.Transfer{}
	row, err := dB.Query("select * from transfer where details like ?", "%"+k+"%")

	//以下为预处理的查询方法
	//stmt, err := dB.Prepare("select * from transfer where details like ?")
	//if err != nil {
	//	fmt.Println(err)
	//	return m, err
	//}
	//defer stmt.Close()
	//row, err := stmt.Query("%" + k + "%")

	//dB.QueryRow("select * from transfer where details like %?%",k)是错误的写法
	//大致意思是：
	//通配符％，应该是参数字符串的一部分，也就是说%必须作为字符串写到参数里面去，而不能在sql语句
	if err != nil {
		fmt.Println("err1:", err)
	}
	defer row.Close()
	for row.Next() {
		i++
		err = row.Scan(&d.Id, &d.UserName, &d.ToWhom, &d.Money, &d.Detail)
		m[i] = d
		if err != nil {
			fmt.Println("err2:", err)
		}
	}
	return m, nil
}

func TransferAddInfo(t model.Transfer) error {
	_, err := dB.Exec("update transfer set details=? where id = ? and whom = ? ", t.Detail, t.Id, t.UserName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
