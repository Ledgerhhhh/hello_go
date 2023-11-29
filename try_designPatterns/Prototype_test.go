package try_designPatterns

import (
	"fmt"
	"testing"
)

// Prototype 接口
type Prototype interface {
	Clone() Prototype
	GetInfo() string
}
type ConcretePrototype struct {
	info string
}

func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		info: c.info,
	}
}

func (c *ConcretePrototype) GetInfo() string {
	return c.info
}

func TestPrototype(t *testing.T) {
	var p Prototype = &ConcretePrototype{
		info: "test",
	}
	p2 := p.Clone()
	fmt.Println(p2.GetInfo())
}
