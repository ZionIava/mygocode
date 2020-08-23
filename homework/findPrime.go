package main
import "fmt"

func findPrimeNum(n int) bool {
	if n == 1 || n == 2 || n == 3{
		return true
	}
	
	for i := 2;i <= n -1 ; i++{
		if n % i == 0{
			return false
		}
	}
	return true
}
func main() {
	var numLimit int 
	var sum int
	var j int = 0
	fmt.Println("请输入你要统计的素数范围")
	fmt.Scanln(&numLimit)
	for i := 1 ; i <= numLimit ;i++ {
		check := findPrimeNum(i)
		if check == true{
			fmt.Printf("%v\t",i)
			j++
			if j % 5 == 0{
				fmt.Println()
			}
			sum += i
		}
	}
	fmt.Println()
	fmt.Println(sum)
}