package main
import "fmt"

func main(){
	var arr = [10]string{"AA","BB","CC","AA","DD","AA","BB","DD","EE","AA"}
	for i := 0 ;i < len(arr); i++{
		if arr[i] == "AA"{
			fmt.Printf("找到了\"AA\",下标为%v\n",i)
		}
	}
}