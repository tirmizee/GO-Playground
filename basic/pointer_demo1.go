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
	fmt.Println("pointerNumber", *pointerNumber, &pointerNumber)

}

func pointerDemo2() {

	x := 10
	fmt.Println("x", x, &x)

	y := x
	fmt.Println("y", y, &y)

	x = 20
	fmt.Println("x", x)
	fmt.Println("y", y)

}
