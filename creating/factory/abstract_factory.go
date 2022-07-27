package factory

import "fmt"

/*
抽象工厂
不公开内部实现的情况下, 让调用者使用你提供的各种功能
*/

type IPerson interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hello! My name is %s, i'm %d years old.\n", p.name, p.age)
}

func NewIPerson(name string, age int) IPerson {
	return person{
		name: name,
		age:  age,
	}
}
