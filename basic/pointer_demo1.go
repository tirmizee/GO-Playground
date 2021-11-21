package main

import (
	"fmt"
)

func pointerDemo() {

	// var number int = 10
	number := 10
	fmt.Println("number", number, &number)

	// var pointerNumber *int = &number
	pointerNumber := &number
	fmt.Println("pointerNumber", *pointerNumber, pointerNumber)

	fmt.Println("after pointer change")
	*pointerNumber = 20
	fmt.Println("number", number, &number)
	fmt.Println("pointerNumber", *pointerNumber, pointerNumber)

}
