package try_designPatterns

import (
	"fmt"
	"testing"
)

// Component 接口
type Component2 interface {
	Operation() string
}

// ConcreteComponent 具体组件
type ConcreteComponent struct{}

func (c ConcreteComponent) Operation() string {
	return "ConcreteComponent operation"
}

type ConcreteComponentDecorator struct {
	component2 Component2
}

func (c ConcreteComponentDecorator) Operation() string {
	return c.component2.Operation() + " Decorator"
}

func Test(t *testing.T) {
	decorator := ConcreteComponentDecorator{
		&ConcreteComponent{},
	}
	fmt.Println(decorator.Operation())
}
