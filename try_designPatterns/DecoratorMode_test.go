package try_designPatterns

import (
	"fmt"
	"testing"
)

type coffee interface {
	Cost() float64
	Description() string
}

type simpleCoffee struct {
}

func (s simpleCoffee) Cost() float64 {
	return 0.2
}

func (s simpleCoffee) Description() string {
	return "存黑卡菲尔"
}

type sugarCoffee struct {
	coffee coffee
}

func (s sugarCoffee) Cost() float64 {
	return s.coffee.Cost() + 0.2
}

func (s sugarCoffee) Description() string {
	return s.coffee.Description() + "加糖"
}

type milkCoffee struct {
	coffee coffee
}

func (m milkCoffee) Cost() float64 {
	return m.coffee.Cost() + 0.2
}

func (m milkCoffee) Description() string {
	return m.coffee.Description() + "加奶"
}

func TestDecorator(t *testing.T) {
	simple := simpleCoffee{}
	m := milkCoffee{
		coffee: simple,
	}
	s := sugarCoffee{
		coffee: simple,
	}
	fmt.Println(m.Cost())
	fmt.Println(m.Description())
	fmt.Println(s.Cost())
	fmt.Println(s.Description())
}
