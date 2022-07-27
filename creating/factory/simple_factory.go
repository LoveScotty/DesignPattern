package factory

import "fmt"

/*
简单工厂模式
简单粗暴, 接受参数, 返回实例
*/

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func (p Person) Greet() {
	fmt.Printf("Hello! My name is %s, i'm %d years old.\n", p.Name, p.Age)
}
