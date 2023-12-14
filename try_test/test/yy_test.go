package test

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"testing"
)

type h struct {
}

func (h h) print(data string) {

}

type HHHH[T int | string] interface {
	print(data T)
}
type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for i := range s {
		sum += s[i]
	}
	return sum
}

type hhhh interface {
	constraints.Ordered
}

func TestH(t *testing.T) {
	var s MySlice[int] = []int{
		11, 12,
	}
	fmt.Println(s.Sum())
}
