package main
import "fmt"


func main() {
	var binaryArr = [5]int{1, 3, 5, 7, 9}
	fmt.Println(binaryArr)
	for i:=0;i<len(binaryArr);i++{
		fmt.Println(binaryArr[len(binaryArr) -1 - int(i)])
	}
}