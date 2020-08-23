package main
import "fmt"

func main(){
	var arrOld = [5]int{1,2,3,4,5}
	var valueAdd int= 6
	var arrNew [6]int
	for i := 0 ;i<len(arrOld);i++{
		arrNew[i]=arrOld[i]
	} 
	arrNew[5]=valueAdd

	
	fmt.Println(arrOld)
	fmt.Println(valueAdd)

	fmt.Println(arrNew)
}