package main
import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now().Unix())
	//fmt.Println(time.Parse("2006-01-02","2018-09-10"))
	var year int 
	var month time.Month
	var day int
	fmt.Println("请输入年份")
	fmt.Scanln(&year)
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	fmt.Println("请输入日份")
	fmt.Scanln(&day)
	changeTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
	oldTime := time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local).Unix()
	timeDiffer := changeTime - oldTime
	//fmt.Println(timeDiffer)
	//fmt.Printf("%v %T %v %T",changeTime,changeTime,oldTime,oldTime)
	var totalDays = timeDiffer / 60 / 60 / 24
	//fmt.Printf("%v %T\n",totalDays,totalDays)
	toDo := totalDays  %  5
	//fmt.Println(toDo)
	if toDo >=0 && toDo <= 2 {
		fmt.Println("打鱼")
	}else {
		fmt.Println("筛网")
	}
}