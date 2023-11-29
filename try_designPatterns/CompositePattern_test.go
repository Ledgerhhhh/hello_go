package try_designPatterns

import (
	"fmt"
	"testing"
)

// Component 接口
type Component interface {
	Operation() string
}

// Leaf 叶子节点
type Leaf struct {
	name string
}

func (l *Leaf) Operation() string {
	return "Leaf " + l.name
}

// Composite 组合节点
type Composite struct {
	name     string
	children []Component
}

func (c *Composite) RemoveNode(Leaf *Leaf) {
	for i := 0; i < len(c.children); i++ {
		if c.children[i] == Leaf {
			c.children = append(c.children[:i], c.children[i+1:]...)
			break
		}
	}
}

func (c *Composite) AddNode() {
	c.children = append(c.children, &Leaf{name: "Leaf A"})
}

func (c *Composite) Operation() string {
	result := "Composite " + c.name + "["
	for _, child := range c.children {
		result += child.Operation() + " "
	}
	result += "]"
	return result
}

func TestComposite(t *testing.T) {
	composite := Composite{}
	composite.children = append(composite.children, &Leaf{name: "Leaf A"})
	composite.children = append(composite.children, &Leaf{name: "Leaf B"})
	composite.children = append(composite.children, &Leaf{name: "Leaf C"})
	fmt.Println(composite.Operation())
}
