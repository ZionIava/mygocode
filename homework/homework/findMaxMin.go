package main
import "fmt"

func findMax(arr [5]int) (int,int){
	//fmt.Println(arr)
	var temp int = 0
	var index int 
	for i := 0;i < len(arr);i++{
		
		if arr[i] > temp {
			temp = arr[i] 
		}
	}
	//fmt.Println(temp)
	//fmt.Println(arr)
	for i := 0;i < len(arr);i++{
		if arr[i] == temp {
			index = i
			break
		}
	}
	return temp,index
}
func findMin(arr [5]int) (int,int) {
	//fmt.Println(arr)
	var temp int = arr[0]
	var index int 
	for i := 0;i < len(arr);i++{
		if temp > arr[i]{
			temp = arr[i]
		}
	}
	//fmt.Println(temp)
	//fmt.Println(arr)
	for i := 0;i < len(arr);i++{
		if arr[i] == temp {
			index = i
			break
		}
	}
	return temp,index
}
func main() {
	var arr [5]int
	for i := 0;i < len(arr); i++{
		fmt.Printf("请输入数组的第%v个元素\n",i+1)
		fmt.Scanln(&arr[i])

	}
	//fmt.Println(arr)
	valueMax,indexMax := findMax(arr)
	fmt.Printf("最大值为%v,其下标为%v\n",valueMax,indexMax)
	valueMin,indexMin := findMin(arr)
	fmt.Printf("最小值为%v,其下标为%v\n",valueMin,indexMin)
}