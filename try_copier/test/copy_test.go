package test

import (
	"fmt"
	"github.com/jinzhu/copier"
	"testing"
)

type User struct {
	name string
	Age  int
	Arr  []int
}

type Employee struct {
	name string
	Age  int
	Arr  []int
	Role string
}

func TestCopy(t *testing.T) {
	user := &User{
		name: "ledger",
		Age:  10,
		Arr:  []int{1, 2, 3, 4, 5, 6},
	}
	emp := &Employee{}
	copier.Copy(emp, user)
	fmt.Printf("%#v\n", emp)
	//test.Employee{name:"", Age:10, Arr:[]int{1, 2, 3, 4, 5, 6}, Role:""}
}

type User2 struct {
	Name string
	Age  int
}

func (u *User2) DoubleAge() int {
	return 2 * u.Age
}

type Employee2 struct {
	Name      string
	DoubleAge int
	Role      string
}

func TestCopyByFromMethod(t *testing.T) {
	user := &User2{
		Name: "ledger",
		Age:  10,
	}
	emp := &Employee2{}
	_ = copier.Copy(emp, user)
	fmt.Printf("%#v\n", emp)
	//test.Employee2{Name:"ledger", DoubleAge:20, Role:""}
}

type User3 struct {
	Name string
	Age  int
	Role string
}
type Employee3 struct {
	Name      string
	Age       int
	SuperRole string
}

func (e *Employee3) Role(role22 string) {
	//e.SuperRole = "Super" + role
	fmt.Println(role22 + "================")
}

func TestCopyByToMethod(t *testing.T) {
	user := User3{Name: "dj", Age: 18, Role: "Admin"}
	employee := Employee3{}

	_ = copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}

type User4 struct {
	Name string
	Age  int
}

type Employee4 struct {
	Name string
	Age  int
	Role string
}

func TestCopyArray(t *testing.T) {
	user4s := []User4{
		{
			Name: "ledger",
			Age:  10,
		},
		{
			Name: "ledger",
			Age:  10,
		},
	}
	var employee4s []Employee4
	_ = copier.Copy(&employee4s, &user4s)
	fmt.Printf("%#v\n", employee4s)
}
func TestAppendArray(t *testing.T) {

	user4 := &User4{
		Name: "ledger",
		Age:  10,
	}
	var employee4s []Employee4
	_ = copier.Copy(&employee4s, user4)
	fmt.Printf("%#v\n", employee4s)
}
