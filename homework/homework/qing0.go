package main
import "fmt"


func main() {
	var binaryArr [3][4]int
	var value int
	for i := 0;i < len(binaryArr);i++{
		for j := 0; j < len(binaryArr[i]);j++{
			fmt.Printf("请输入第%v行第%v列的值\n",i+1,j+1)
			fmt.Scanln(&value)
			binaryArr[i][j]=value
		}
	}
	//fmt.Println(binaryArr)
	for i := 0; i < len(binaryArr);i++{
		for j := 0;j < len(binaryArr[i]);j++{
			if i == 0 || i == 2 || j == 0 || j == 3{
				fmt.Printf("0\t")
			}else {
				fmt.Printf("%v\t",binaryArr[i][j])
			}
		}
		fmt.Println()
		fmt.Println()
	}
}