package main

import "fmt"

// operators

func operatorsDemo()  {

	const A = 34
	const B = 27

	addition := A + B
	subtraction := A - B
	multiplication := A * B
	division := A / B
	mod := A % B

	fmt.Println(addition)
	fmt.Println(subtraction)
	fmt.Println(multiplication)
	fmt.Println(division)
	fmt.Println(mod)

}