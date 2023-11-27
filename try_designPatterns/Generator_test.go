package try_designPatterns

import (
	"fmt"
	"testing"
)

type Computer struct {
	CPU    string
	Memory string
	Disk   string
}

// Builder 生成器接口
type ComputerBuilder interface {
	BuildCPU()
	BuildMemory()
	BuildDisk()
	GetComputer() Computer
}

type ConcreteBuilder struct {
	computer Computer
}

func (c *ConcreteBuilder) BuildCPU() {
	c.computer.CPU = "Intel Core i7"
}

func (c *ConcreteBuilder) BuildMemory() {
	c.computer.Memory = "16GB DDR4"
}

func (c *ConcreteBuilder) BuildDisk() {
	c.computer.Disk = "1TB SSD"
}

func (c *ConcreteBuilder) GetComputer() Computer {
	return c.computer
}

type Director struct {
	builder ComputerBuilder
}

func (receiver *Director) GetPc() Computer {
	receiver.builder.BuildCPU()
	receiver.builder.BuildMemory()
	receiver.builder.BuildDisk()
	return receiver.builder.GetComputer()
}

func TestGenerator(t *testing.T) {
	builder := ConcreteBuilder{}
	var computerBuilder = builder
	pc := (&Director{builder: &computerBuilder}).GetPc()
	fmt.Printf("%+v", pc)
}
