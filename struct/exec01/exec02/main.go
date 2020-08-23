package main
import "fmt"

type twoDimenArr struct{
	arr [3][3]int	
}

func (tDA twoDimenArr) switchArr() ([3][3]int,[3][3]int) {
	for i:=0;i< len(tDA.arr);i++{
		for j :=0;j<len(tDA.arr[i]);j++{
			fmt.Printf("请输入第%v行第%v列的值:\n",i+1,j+1)
			fmt.Scanln(&tDA.arr[i][j])
		}
	}
	var arrNew [3][3]int
	for i:=0;i< len(tDA.arr);i++{	
		for j :=0;j < len(tDA.arr[i]);j++{
			if i != j {
				arrNew[j][i] = tDA.arr[i][j]
			}else {
				arrNew[i][j] = tDA.arr[i][j]
			}
		}	
		fmt.Println()
	}
	return tDA.arr,arrNew
}


func main() {
	var abc twoDimenArr
	fmt.Printf("%T",abc)
	defOld,defNew := abc.switchArr()

	for i := 0;i<len(defOld);i++{
		for j:=0; j<len(defOld[i]);j++{
			fmt.Printf("%v\t",defOld[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	for i := 0;i<len(defNew);i++{
		for j:=0; j<len(defNew[i]);j++{
			fmt.Printf("%v\t",defNew[i][j])
		}
		fmt.Println()
	}


}