package try_designPatterns

import (
	"fmt"
	"testing"
)

// Flyweight 享元接口
type Flyweight interface {
	PrintCharacter(font, color string)
}

// ConcreteFlyweight 具体享元对象
type ConcreteFlyweight struct {
	character string
}

// PrintCharacter 实现 Flyweight 接口的方法
func (f *ConcreteFlyweight) PrintCharacter(font, color string) {
	fmt.Printf("Character: %s, Font: %s, Color: %s\n", f.character, font, color)
}

// FlyweightFactory 享元工厂
type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}
}

// GetFlyweight 获取享元对象
func (factory *FlyweightFactory) GetFlyweight(character string) Flyweight {
	if flyweight, ok := factory.flyweights[character]; ok {
		return flyweight
	}

	// 如果不存在，则创建新的享元对象并存储起来
	flyweight := &ConcreteFlyweight{character: character}
	factory.flyweights[character] = flyweight
	return flyweight
}

func TestFlyweight(t *testing.T) {
	factory := NewFlyweightFactory()

	// 获取并共享享元对象
	flyweightA := factory.GetFlyweight("A")
	flyweightB := factory.GetFlyweight("B")
	flyweightC := factory.GetFlyweight("A")

	// 通过享元对象打印字符
	flyweightA.PrintCharacter("FontA", "ColorA")
	flyweightB.PrintCharacter("FontB", "ColorB")
	flyweightC.PrintCharacter("FontC", "ColorC")
}
