package try_designPatterns

import (
	"fmt"
	"testing"
)

type System interface {
	turnOn()
	turnOff()
}

// SubsystemA 子系统A
type SubsystemA struct{}

func (s SubsystemA) turnOn() {
	fmt.Println("SubsystemA: Turning on")
}

func (s SubsystemA) turnOff() {
	fmt.Println("SubsystemA: Turning off")
}

type SubsystemB struct{}

func (s SubsystemB) turnOn() {
	fmt.Println("SubsystemB: Turning on")
}

func (s SubsystemB) turnOff() {
	fmt.Println("SubsystemB: Turning off")
}

// SubsystemC 子系统C
type SubsystemC struct{}

func (s SubsystemC) turnOn() {
	fmt.Println("SubsystemC: Turning on")
}

func (s SubsystemC) turnOff() {
	fmt.Println("SubsystemC: Turning off")
}

// Facade 外观类
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
	subsystemC *SubsystemC
}

func NewFacade(subsystemA *SubsystemA, subsystemB *SubsystemB, subsystemC *SubsystemC) *Facade {
	return &Facade{subsystemA: subsystemA, subsystemB: subsystemB, subsystemC: subsystemC}
}

// TurnOn 打开音响系统
func (f *Facade) TurnOn() {
	fmt.Println("Facade: Turning on the entertainment system")
	f.subsystemA.turnOn()
	f.subsystemB.turnOn()
	f.subsystemC.turnOn()
}

// TurnOff 关闭音响系统
func (f *Facade) TurnOff() {
	fmt.Println("Facade: Turning off the entertainment system")
	f.subsystemA.turnOff()
	f.subsystemB.turnOff()
	f.subsystemC.turnOff()
}

func Test2(t *testing.T) {
	facade := NewFacade(&SubsystemA{}, &SubsystemB{}, &SubsystemC{})
	facade.TurnOn()
	facade.TurnOff()
}
