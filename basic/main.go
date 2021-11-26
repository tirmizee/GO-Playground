package main

import (
	"fmt"
)

type EmptyInterface interface{}

func main() {

	var emptyInterface EmptyInterface

	emptyInterface = 1
	fmt.Printf("%T %v \n", emptyInterface, emptyInterface)

	emptyInterface = 1.12
	fmt.Printf("%T %v \n", emptyInterface, emptyInterface)

	emptyInterface = true
	fmt.Printf("%T %v \n", emptyInterface, emptyInterface)

	emptyInterface = "Hello"
	fmt.Printf("%T %v \n", emptyInterface, emptyInterface)

	emptyInterface = Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}
	fmt.Printf("%T %v \n", emptyInterface, emptyInterface)

}

type SalaryCalculator interface {
	CalculateSalary() int
}

type NameCalculator interface {
	GetName() string
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (p Permanent) GetName() string {
	return "anonymous"
}

type Contract struct {
	empId    int
	basicpay int
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func printSalary(s SalaryCalculator) {
	fmt.Println(s.CalculateSalary())
}
