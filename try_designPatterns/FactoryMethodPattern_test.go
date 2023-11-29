package try_designPatterns

import (
	"fmt"
	"testing"
)

type product interface {
	getName() string
}
type productA struct {
}
type productB struct {
}

func (p *productA) getName() string {
	return "Concrete Product A"
}

func (p *productB) getName() string {
	return "Concrete Product B"
}

type Creator interface {
	getProduct() product
}

type productAFactory struct {
}

func (p *productAFactory) getProduct() product {
	return &productA{}
}

type productBFactory struct {
}

func (p *productBFactory) getProduct() product {
	return &productB{}
}

func TestFactoryMethodPattern_test(t *testing.T) {
	factory := productAFactory{}
	fmt.Println(factory.getProduct().getName())
	bFactory := productBFactory{}
	fmt.Println(bFactory.getProduct().getName())
}
