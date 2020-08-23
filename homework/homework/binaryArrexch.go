package main
import "fmt"

func exchangeNum(arrOld [4][4]int) (arrNew [4][4]int)  {
	//fmt.Println(len(arrOld))
	//var middle [4][4]int
	for i := 0;i < len(arrOld) ;i++{
		//fmt.Println(arrOld[i])
		//fmt.Println(arrOld[i])
		//fmt.Println()
		//fmt.Printf("%v %T\n",arrOld[i],arrOld[i])
		//fmt.Printf("%v %T\n",len(arrOld[i])-1,int(len(arrOld[i])))		
		//arrOld[i] , arrNew[len(arrOld) - int(i)] = arrNew[len(arrOld) - int(i)] ,arrOld[i]
		arrNew[len(arrNew) - 1 -int(i)] = arrOld[i]  
	}
	//fmt.Println(arrNew)
	return arrNew
}
func PrintArr(arr [4][4]int) {
	for i:=0 ;i<len(arr);i++{
		for j:=0 ;j < len(arr[i]);j++{
			fmt.Printf("%v\t",arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
func main() {
	var binaryArr [4][4]int
	var value int
	var binaryNew [4][4]int
	for i := 0;i < len(binaryArr); i++ {
		for j := 0;j < len(binaryArr[i]);j++{
			fmt.Printf("请输入第%v行第%v列的值\n",i+1 , j+1)
			fmt.Scanln(&value)
			binaryArr[i][j] = value
		}
	}
	//fmt.Println(binaryArr)
	//fmt.Println(binaryArr[1])
	//fmt.Println(binaryArr[2])
	//fmt.Println(binaryArr[3])
	binaryNew=exchangeNum(binaryArr)
	//fmt.Println(binaryNew)
	PrintArr(binaryArr)
	PrintArr(binaryNew)

}