package main
import "fmt"

func main() {
	var Sum float64= 82
	var mus float64= 5
	avg := fmt.Sprintf("%.2f", float64(Sum) / float64(mus))
	fmt.Printf("%T",avg)
}

