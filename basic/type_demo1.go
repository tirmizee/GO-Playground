package main

import "fmt"

type ListInt []int
type ListBool []bool
type ListString []string

func (l ListInt) Print() {
	for _, n := range l {
		fmt.Print(n, " ")
	} 
	fmt.Println()
}

func (l ListBool) Print() {
	for _, n := range l {
		fmt.Print(n, " ")
	} 
	fmt.Println()
}

func (l ListString) Print() {
	for _, n := range l {
		fmt.Print(n, " ")
	} 
	fmt.Println()
}

func demoType11()  {

	numbers := ListInt{1,2,3,4}
	numbers.Print()

	var names = ListString{"Java", "Kotlin", "Go"}
	names.Print()

	var statuses ListBool = ListBool{true, false}
	statuses.Print()

}