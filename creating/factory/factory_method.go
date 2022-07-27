package factory

/*
工厂方法模式
通过实现工厂函数来创建多种工厂, 将对象创建从由一个对象负责所有具体类的实例化, 变成由一群子类来负责对具体类的实例化
*/

func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			Name: name,
			Age:  age,
		}
	}
}
