package try_designPatterns

import (
	"fmt"
	"testing"
)

type Animal interface {
	Speak()
}
type dog struct {
}

type cat struct {
}

func (c cat) Speak() {
	fmt.Println("喵喵喵")
}

func (d dog) Speak() {
	fmt.Println("汪汪汪")
}

func AnimalFactory(t string) Animal {
	switch t {
	case "dog":
		return &dog{}
	case "cat":
		return &cat{}
	default:
		return nil
	}
}

func TestFactory(t *testing.T) {
	AnimalFactory("dog").Speak()
}
