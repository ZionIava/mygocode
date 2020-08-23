package main

import (
	"fmt"
	"gocode/homework/encapsulation/exec01/model"
)


func main() {
	acc := model.Account(01)
	acc.SetAccount("1522507")
	acc.SetBalance(6554.5)
	acc.SetPassword("123456")
	fmt.Println(acc.ID,acc.GetAccount(),acc.GetBalance(),acc.GetPwd())
}