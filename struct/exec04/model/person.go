package model
import "fmt"


type personEncaps struct {
	Name string
	age uint
	salt float64
}

func Person(name string) *personEncaps {
	return &personEncaps{
		Name : name,
	}
} 

func (p *personEncaps) SetAge(n uint)  {
	if n > 0 && n < 150 {
		p.age = n
	}else {
		fmt.Println("年龄输入有误")
	}
}

func (p *personEncaps) GetAge() uint {
	return p.age
}

func (p *personEncaps) SetSalt(n float64) {
	if n >=3000 && n <= 30000{
		p.salt = n
	}else {
		fmt.Println("工资输入有误")
	}
}

func (p *personEncaps) GetSalt() float64 {
	return p.salt
}

