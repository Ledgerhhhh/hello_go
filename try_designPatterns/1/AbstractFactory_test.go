package try_designPatterns

import "testing"

type Shape interface {
	Draw() string
}

type Circle struct{}

func (c Circle) Draw() string {
	return "画一个圆形"
}

type Color interface {
	Fill() string
}

type Red struct{}

func (r Red) Fill() string {
	return "填充红色"
}

type AbstractFactory interface {
	CreateShape() Shape
	CreateColor() Color
}

type ConcreteFactory struct{}

func (c ConcreteFactory) CreateShape() Shape {
	return &Circle{}
}

func (c ConcreteFactory) CreateColor() Color {
	return &Red{}
}

func TestAbstractFactory(t *testing.T) {
	c := &ConcreteFactory{}
	println(c.CreateShape().Draw())
	println(c.CreateColor().Fill())
}
