package model
import "fmt"


type account struct{
	ID int
	account string
	balance float64
	password string
}

func Account(id int) *account {
	return &account {
		ID : id ,
	}
}

func (id *account) SetAccount(acc string) {
	if len(acc) >=6 && len(acc) <= 10{
		id.account = acc
	}else {
		fmt.Println("账号长度异常")
	}
}

func (id *account) SetBalance(balance float64) {
	if balance > 20 {
		id.balance = balance
	}else {
		fmt.Println("余额输入有误")
	}
}

func (id *account) SetPassword(pwd string) {
	if len(pwd) == 6 {
		id.password = pwd
	}else {
		fmt.Println("密码位数不对，请重新设置")
	}
}

func (id *account) GetAccount() string {
	return id.account
}

func (id *account) GetBalance() float64 {
	return id.balance
}

func (id * account) GetPwd() string {
	return id.password
}