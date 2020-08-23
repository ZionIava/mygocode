package main
import "fmt"


func main() {

		var (
			year int
			month int
			day int
			check string
		)
		for{
			lable_init:
			fmt.Println("请输入年份")
			fmt.Scanln(&year)
			if year <= 0 {
				fmt.Println("年份不正确")
				goto lable_init
			} 
			lable_month:
			fmt.Println("请输入月份")
			fmt.Scanln(&month)
			if month <= 0 || month > 12 {
				fmt.Println("月份不正确")
				goto lable_month
			}
			lable_day:
			fmt.Println("请输入日")
			fmt.Scanln(&day)
			
			switch month {
				case 2:
					if (year % 4 ==0 && year % 100 != 0) || year % 400 == 0{
						if day > 29 || day <= 0 {
							fmt.Println("日数错误")
							goto lable_day
						}else {
							goto lable_print
						}
					}else {
						if day > 28 ||day <=0 {
							fmt.Println("日数错误")
							goto lable_day
						}else {
							goto lable_print
						}
					}
				case 1,3,5,7,8,10,12:
					if day > 31|| day <= 0 {
						fmt.Println("日数错误")
						goto lable_day
					}else {
						goto lable_print
					}
				case 4,6,9,11:
					if day > 30 || day <= 0{
						fmt.Println("日数错误")
						goto lable_day
					}else {
						goto lable_print
					}
				default:
					fmt.Println("月份错误")
					goto lable_month
			}	
			lable_print:
			fmt.Printf("%v年%v月%v日\n",year,month,day)
			fmt.Println("是否继续:yes/no")
			fmt.Scanln(&check)
			if check =="yes"{
				goto lable_init
			}else {
				break
			}
		}
			
		

}