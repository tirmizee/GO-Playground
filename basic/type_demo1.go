package main

import "fmt"

type ListInt []int
type ListBool []bool
type ListString []string

func demoType1()  {

	numbers := ListInt{1,2,3,4}
	fmt.Println(numbers)

	var names = ListString{"Java", "Kotlin", "Go"}
	fmt.Println(names)

	var statuses ListBool = ListBool{true, false}
	fmt.Println(statuses)

}