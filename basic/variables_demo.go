package main

import "fmt"

// declare variable with manual

func declareVariableDemo() {

	var text string 
	var number int
	var number8 int8
	var number16 int16
	var isActive bool
	var amount float64
	
	text = "String type"
	number = 129
	number8 = 127
	number16 = -32768
	isActive = true
	amount = 20000.89

	fmt.Println(text)
	fmt.Println(number)
	fmt.Println(number8)
	fmt.Println(number16)
	fmt.Println(isActive)
	fmt.Println(amount)

}

// declare variable with default

func declareVariableWithDefaultDemo() {

	var text string = "String type"
	fmt.Println(text)

	var number int = 129
	fmt.Println(number)

	var isActive bool = true
	fmt.Println(isActive)

	var amount float64 = 20000.89
	fmt.Println(amount)

}

// declare variable auto detect type with var

func declareVariableWithDefaultDemo2() {

	var text = "String type"
	fmt.Println(text)

	var number = 129
	fmt.Println(number)

	var isActive = true
	fmt.Println(isActive)

	var amount = 20000.89
	fmt.Println(amount)

}

// declare variable auto detect type with :=

func declareVariableAutoDemo() {

	text := "String type"
	fmt.Println(text)

	number := 129
	fmt.Println(number)

	isActive := true
	fmt.Println(isActive)

	amount := 20000.89
	fmt.Println(amount)

}
