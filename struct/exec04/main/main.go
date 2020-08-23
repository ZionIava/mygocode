package main
import (
	"fmt"
	"gocode/struct/exec04/model"
)


func main() {
	person := model.Person("tom")
	person.SetAge(18)
	person.SetSalt(7000)
	fmt.Println(*person)
	fmt.Println(person.Name,person.GetAge(),person.GetSalt())

	
}