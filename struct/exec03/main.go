package  main
import "fmt"

type account struct{
		Account int
		Password string
		Save float64
}

func (a *account) saveMoney() {
	var (
		Account int
		Password string
		num float64
	)
	fmt.Println("请输入账号")
	fmt.Scanln(&Account)
	fmt.Println("请输入密码")
	fmt.Scanln(&Password)
	if Account == a.Account && Password == a.Password{
		fmt.Println("请输入存款金额，并将纸币放入存钞口")
		fmt.Scanln(&num)
		a.Save += num
		fmt.Printf("存入%v,存款成功，当前余额为%v\n",num,a.Save)
	}else {
		fmt.Println("账号和密码不匹配，请检查账号和密码的输入")
		return
	}
	
}

func (a *account) getMoney() {
	var (
		Account int
		Password string
		num float64
	)
	fmt.Println("请输入账号")
	fmt.Scanln(&Account)
	fmt.Println("请输入密码")
	fmt.Scanln(&Password)
	if Account == a.Account && Password == a.Password{
		fmt.Println("请输入取款金额")
		fmt.Scanln(&num)
		//fmt.Println(num)
		if num > 0 && num <= a.Save {
			a.Save -= num
		fmt.Printf("取款%v,存款成功，当前余额为%v\n",num,a.Save)
		}else {
			fmt.Println("取款金额不对，请重新选择业务类型")
		}
	}else {
		fmt.Println("账号和密码不匹配，请检查账号和密码的输入")
		return
	}
	
}

func (a *account) queryMoney() {
	var (
		Account int
		Password string
	)
	fmt.Println("请输入账号")
	fmt.Scanln(&Account)
	fmt.Println("请输入密码")
	fmt.Scanln(&Password)
	if Account == a.Account && Password == a.Password{
		fmt.Printf("当前余额为%v\n",a.Save)

	}else {
		fmt.Println("账号和密码不匹配，请检查账号和密码的输入")
		return
	}
	
}



func main() {

	myBussness := account{
		Account : 123321,
		Password :"666666",
		Save : 100,
	} 
	var bussnessType int
	var ifGoOn string
	busschoice:
	fmt.Println("请选择你要办理的业务并输入对应业务代码")
	fmt.Printf("1:查询；\n2:存款；\n3:取款\n")
	fmt.Scanln(&bussnessType)
	
	
	for {
		if bussnessType == 1 {
			myBussness.queryMoney()
		}else if bussnessType == 2 {
			myBussness.saveMoney()
		}else if bussnessType == 3 {
			myBussness.getMoney()
		}else {
			fmt.Println("请核对你要办理的业务以及输入的业务代码")
			return
		}
		fmt.Println("是否继续办理业务")
		fmt.Println("yes/no")
		fmt.Scanln(&ifGoOn)
		if ifGoOn == "no"{
			break
		}else{
			goto busschoice
		}
	}
	
}