package main
import "fmt"

type MethodUtils struct {
	n int
}

func (method *MethodUtils) cfb() {
	
	for i := 1;i <= method.n;i++{
		for j := 1;j<= i;j++{
			fmt.Printf("%v X %v = %v\t", j,i,i*j)
		}
		fmt.Println()
	}
}

func main() {
	var method MethodUtils
	var n int
	var a *string
	carryon:
	fmt.Println("请输入乘数")
	fmt.Scanln(&n)
	method.n = n
	
	method.cfb()
	fmt.Printf("按任意键继续...")
	fmt.Scanln(&a)

	if *a !=""  {
		goto carryon
	}
	

}