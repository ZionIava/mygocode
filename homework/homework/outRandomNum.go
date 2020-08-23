package main
import (
	"time"
	"fmt"
	"math/rand"
)
func outputAMM(arr [10]int) (avg int,maxNum int ,minNum int)  {
	maxNum = 0
	sum := 0
	avg = 0
	//fmt.Println(arr)
	for i := 0; i < len(arr); i++{
		if arr[i] > maxNum {

			maxNum = arr[i]
		}
		//if arr[i] > arr[i+1]{
		//	minNum = arr[i + 1]
		//}else {
		//	minNum = arr[i]
		//}
		sum += arr[i]

		avg = sum / len(arr)	
	//fmt.Println(arr[i])
	}
	minNum = maxNum
	for i := 0;i <len(arr);i++{
		if arr[i]< minNum{
			minNum,arr[i] = arr[i], minNum
		}
	}
	
	return avg  , maxNum , minNum
}
	
func main() {
	var arr [10]int
	for i := 0 ; i < 10 ; i++{
		time.Sleep(time.Duration(1)*time.Second)
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	avg,maxNum,minNum:=outputAMM(arr)
	fmt.Printf("平均值为：%v\t最大值为%v\t最小值为%v\n",avg,maxNum,minNum)
}