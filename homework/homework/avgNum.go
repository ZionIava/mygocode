package main
import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)
func countNum(arr [10]float64)(float64,float64,float64,float64){
	var (
		Sum float64
		bigNum float64
		smallNum float64
		equalNum float64
	) 
	for i := 0 ;i < len(arr);i++{
		Sum += arr[i]
	}
	avg1 := fmt.Sprintf("%.2f", float64(Sum) / float64(len(arr)))
	avg ,_:= strconv.ParseFloat(avg1,64) 
	//avg = fmt.Sprintf("%.2f",avg1)
	for i := 0;i < len(arr); i++{
		if arr[i] > avg {
			bigNum++
		}else if arr[i] < avg {
			smallNum++
		}else {
			equalNum++
		}
	}
	return avg,bigNum,smallNum,equalNum
}


func main() {
	var arr [10]float64
	for i := 0;i < len(arr); i++{
		time.Sleep(time.Duration(1)*time.Millisecond)
		rand.Seed(time.Now().UnixNano())	
		arr[i] = float64(rand.Intn(100))
		//fmt.Println("请输入第%v个元素的值",i + 1)
		//t.Scanln(&arr[i])
	}
	fmt.Println(arr)
	avg,bigNum,smallNum,equalNum := countNum(arr)
	fmt.Printf("比平均数%v大的有%v个\n",avg,bigNum)
	fmt.Printf("比平均数%v小的有%v个\n",avg,smallNum)
	fmt.Printf("等于平均数%v的有%v个\n",avg,equalNum)

}