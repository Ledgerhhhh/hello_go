package main

import (
	"fmt"
	"testing"
)

type ss1 struct {
	Name    string
	Age     int
	Address string
}
type ss2 struct {
	Id    string
	Email string
}
type ssi interface {
	ss1 | ss2
}

func f[T ssi](ssi T) {
	fmt.Printf("%+v", ssi)
}
func TestK(t *testing.T) {
	s := ss1{
		Name:    "hhh",
		Age:     1623,
		Address: "china",
	}
	s3 := ss2{
		Id:    "151561",
		Email: "215125",
	}
	f(s)
	f(s3)
}
