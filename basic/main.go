package main

import "fmt"

const PI = 3.14

func main()  {
	operatorsDemo()
}

func printDemo() {
	fmt.Println(PI)
	fmt.Println("Hello GO")
	fmt.Println("Hello", "GO")
	fmt.Println("Hello", "GO", 2564)
	fmt.Println(1994, 2564)
}

func constDemo() {

	const NAME = "Pratya Yeekhaday"
	const AGE = 27
	const SINGLE = true

	fmt.Println(NAME)
	fmt.Println(AGE)
	fmt.Println(SINGLE)

}

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
