package main
import (
	"fmt"
	"time"
	"math/rand"
)

func BubbleSort(arr *[10]int) (*[10]int) {
	var checkNum int = 0
	for i := 0 ; i < len(*arr); i++ {
		for j := 0 ; j < len(*arr) -1 -i;j++{
			if (*arr)[j] > (*arr)[j+1]{
				checkNum = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = checkNum
			}
		}
	}
	return arr
}

func binaryCheck(arr *[10]int, leftIndex int,rightIndex int,findVal int) {
	
	//var returnIndex int
	if leftIndex > rightIndex{
		fmt.Printf("找不到%v\n",findVal)
		//return 0,false
	}else {
		var middleIndex = (leftIndex + rightIndex) / 2
		
		if findVal < (*arr)[middleIndex]{
			binaryCheck(arr  , leftIndex,middleIndex-1,findVal)
		}else if findVal > (*arr)[middleIndex] {
			binaryCheck(arr  ,middleIndex +1 ,rightIndex,findVal)
		}else {
			
			fmt.Println(middleIndex)
			
		}
		//return middleIndex, true
	}
	
}

func main() {
	var arr [10]int
	for i := 0 ;i < 10;i++{
		time.Sleep(time.Duration(100)*time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	BubbleSort(&arr)
	binaryCheck(&arr,0,9,90)
	//mt.Printf("%v %t\n",indexNum,rst)
	//f rst == true {
	//	fmt.Printf("找到了，下标为%v\n",indexNum)
	//
	fmt.Println(arr)

}