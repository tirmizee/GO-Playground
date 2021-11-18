package main

import "fmt"

const PI = 3.14

func main()  {
	sliceArraysDemo()
}

func printDemo() {
	fmt.Println(PI)
	fmt.Println("Hello GO")
	fmt.Println("Hello", "GO")
	fmt.Println("Hello", "GO", 2564)
	fmt.Println(1994, 2564)
}

// declare constant variable

func constDemo() {

	const NAME = "Pratya Yeekhaday"
	const AGE = 27
	const SINGLE = true

	fmt.Println(NAME)
	fmt.Println(AGE)
	fmt.Println(SINGLE)

}

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

// declare variable auto detect type

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

// declare variable auto detect type

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

// arrays 

func arraysDemo()  {

	var students1 [3]string
	students1[0] = "students1 a"
	students1[1] = "students1 b"
	students1[2] = "students1 c"
	fmt.Println(students1[0])
	fmt.Println(students1[1])
	fmt.Println(students1[2])


	var students2 [3]string = [3]string {
		"students2 a",
		"students2 b",
		"students2 c",
	}
	fmt.Println(students2[0])
	fmt.Println(students2[1])
	fmt.Println(students2[2])

	var students3 []string = []string {
		"students3 a",
		"students3 b",
		"students3 c",
	}
	fmt.Println(students3[0])
	fmt.Println(students3[1])
	fmt.Println(students3[2])

	students4 := [3]string {
		"students4 a",
		"students4 b",
		"students4 c",
	}
	fmt.Println(students4[0])
	fmt.Println(students4[1])
	fmt.Println(students4[2])

	students5 := []string {
		"students5 a",
		"students5 b",
		"students5 c",
	}
	fmt.Println(students5[0])
	fmt.Println(students5[1])
	fmt.Println(students5[2])

	students6 := make([]string, 3)
	students6[0] = "students6 a"
	students6[1] = "students6 b"
	students6[2] = "students6 c"
	fmt.Println(students6)

}

// slice arrays 

func sliceArraysDemo()  {

	numbers := make([]int, 3)
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	fmt.Println(numbers)

	numbers2 := numbers[:3]
	fmt.Println(numbers2)

	numbers3 := numbers[1:3]
	fmt.Println(numbers3)

	numbers4 := numbers[2:3]
	fmt.Println(numbers4)

	numbers5 := numbers[2:]
	fmt.Println(numbers5)

	numbers6 := numbers[1:]
	fmt.Println(numbers6)

}

// custom type

type desk []string


// func with return 

func funcWithReturnDemo()  {

	newText := newText()
	newNumber := newNumber()
	isActive := isActive()

	fmt.Println(newText)
	fmt.Println(newNumber)
	fmt.Println(isActive)
}

func newText() string {
	return "New Card"
}

func newNumber() int {
	return 555
}

func isActive() bool {
	return 27 > 5
}


// func with parameter

func funcWithParameterDemo()  {
	fmt.Println(say("tirmizee"))
	fmt.Println(add(5, 10))
}

func say(name string) string {
	return "Hello " + name
}

func add(num1 int, num2 int) int {
	return num1 + num2
}