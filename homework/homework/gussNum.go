package main
import (
	"fmt"
	"time"
	"math/rand"
)
func main(){
	rand.Seed(time.Now().Unix())
	randomNum := rand.Intn(100)
	fmt.Println(randomNum)
	var gussNum int
	var i int
	for i = 1 ;i <= 10; i++{
		//lable_init:
		fmt.Println("请在1-100中猜一个你数")
		fmt.Scanln(&gussNum)
		//fmt.Printf("%v %T %v %T\n",randomNum,randomNum,gussNum,gussNum)
		if gussNum == randomNum {
			switch i {
				case 1:					
					fmt.Println("天才")
				case 2,3 :					
					fmt.Println("厉害")
				case 4,5,6,7,8,9:					
					fmt.Println("还行")
				case 10:
					fmt.Println("终于")
				default :
					fmt.Println("程序貌似出错了")
			}
			break
		}
	}
	if i == 11{
	fmt.Println("10次都猜不到吗？")	
	}
}

