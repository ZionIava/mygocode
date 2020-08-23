package main
import "fmt"
type cat struct {
	Name string
	Age int
	Color string
}


func main() {

	var cat1 cat 
	cat1.Name = "小白"
	cat1.Age = 1
	cat1.Color = "白色"
	fmt.Printf("%p\n",&cat1.Name)
	fmt.Printf("%p\n",&cat1.Age)
	fmt.Printf("%p\n",&cat1.Color)
	fmt.Printf("%p\n",&cat1)
	fmt.Println()
	var c3 *cat = new(cat)
	//c3.Name = "abc"
	//c3.Age= 15
	//c3.Color = "b"
	fmt.Println(c3)
	fmt.Printf("%T\n",c3)
	fmt.Println(*c3)
	fmt.Printf("%p\n",c3)
	fmt.Printf("%p",&c3)

	fmt.Println("=============")

	var c4 *cat = &cat {"def",2,"w"}
	fmt.Printf("%T\n",c4)
	fmt.Printf("%p\n",c4)
	fmt.Printf("%v\n",c4)
	fmt.Printf("%v\n",*c4)
	fmt.Printf("%v\n",&c4)
	
}