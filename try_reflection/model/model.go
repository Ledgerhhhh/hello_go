package model

import "fmt"

//type pi interface {
//	PP()
//	pp()
//}

type hh struct {
}

type P struct {
	Name string
	age  int
	hh   hh
}
type P2 struct {
	Name string `json:"name" xml:"name"`
	age  int    `json:"age" xml:"age"`
}

func NewP2(name string, age int) *P2 {
	return &P2{Name: name, age: age}
}

func (p *P) Say(words string) string {
	fmt.Println(words)
	return words
}

func NewP(name string, age int) *P {
	return &P{Name: name, age: age}
}

func (receiver *P) PP() {
	fmt.Println("PP")
}
func (receiver *P) pp() {
	fmt.Println("pp")
}

// Employee 结构体表示一个雇员
type Employee struct {
	ID       int
	Name     string `json:"name" xml:"name"`
	Age      int
	Salary   float64
	Location string
	Slice    []int
	Array    [2]int
}

func NewEmployee(ID int, name string, age int, salary float64, location string) *Employee {
	e := &Employee{ID: ID, Name: name, Age: age, Salary: salary, Location: location}
	var ints []int
	var ints2 [2]int
	ints = append(ints, 10)
	ints2[0] = 11
	ints2[1] = 11
	e.Slice = ints
	e.Array = ints2
	return e

}

// 方法1：计算雇员的年薪
func (e *Employee) AnnualSalary() float64 {
	return e.Salary * 12
}

// 方法2：更新雇员的位置
func (e *Employee) UpdateLocation(newLocation string) {
	e.Location = newLocation
}

// 方法3：打印雇员的信息
func (e *Employee) PrintInfo() {
	fmt.Printf("ID: %d, Name: %s, Age: %d, Salary: %.2f, Location: %s\n", e.ID, e.Name, e.Age, e.Salary, e.Location)
}

// 方法4：根据给定百分比增加雇员的薪水
func (e *Employee) IncreaseSalaryByPercentage(percentage float64) {
	e.Salary = e.Salary * (1 + percentage/100)
}
