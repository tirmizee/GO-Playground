package main

import "fmt"

const PI = 3.14
const NAME = "Pratya Yeekhaday"
const AGE = 27
const SINGLE = true

// grouping constant

const (
	ONE = 1
	TWO = 2
	THREE = 3
)

func constDemo() {
	fmt.Println(PI)
	fmt.Println(NAME)
	fmt.Println(AGE)
	fmt.Println(SINGLE)

	fmt.Println(ONE)
	fmt.Println(TWO)
	fmt.Println(THREE)
}
