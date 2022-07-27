package factory

import "testing"

func TestSimpleFactory(t *testing.T) {
	p := NewPerson("scotty", 24)
	p.Greet()
}

func TestAbstractFactory(t *testing.T) {
	p := NewIPerson("scotty", 24)
	p.Greet()
}

func TestFactoryMethod(t *testing.T) {
	p := NewPersonFactory(24)
	a := p("scotty")
	a.Greet()
	
	b := p("scott")
	b.Greet()
}
