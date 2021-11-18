package main

import "fmt"


// func with return 

func funcWithReturnDemo() {

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