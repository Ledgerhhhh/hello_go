package test

import (
	"fmt"
	"testing"
)

type mys1 struct {
	name string
}

func (m mys1) call() {
	//TODO implement me
	panic("implement me")
}

type mys2 struct {
	name2 string
}

type fff interface {
	call()
}

func GetAllFiled[T fff](hh []T) {
	for i := range hh {
		fmt.Println(hh[i])
	}
}
func TestG(t *testing.T) {
	m := []mys1{
		{
			name: "ledger",
		},
	}
	GetAllFiled[mys1](m)
}
