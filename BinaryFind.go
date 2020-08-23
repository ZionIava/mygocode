package main
import "fmt"

func BinaryFind(arr *[6]int, leftIndex int ,rightIndex int , findVal int) {
	
	if leftIndex > rightIndex{
		fmt.Printf("找不到%v\n",findVal)
		return
	}
	fmt.Printf("%T %v\n",arr,arr)
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal{
		BinaryFind(arr, leftIndex, middle -1,findVal)
	}else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle + 1, rightIndex, findVal)
	}else {
		fmt.Printf("找到了，%v的下标为%v\n",findVal,middle)
	}
}

func main() {
	var arr = [6]int{1, 8, 10, 89, 100, 1001}
	BinaryFind(&arr , 0 , len(arr) - 1, 1001)
}