package main

import "fmt"

func funcWithReturnDemo() {

	newText := newText()
	newNumber := newNumber()
	isActive := isActive()

	fmt.Println(newText)
	fmt.Println(newNumber)
	fmt.Println(isActive)
}

// func with return 

func newText() string {
	return "New Card"
}

func newNumber() int {
	return 555
}

func isActive() bool {
	return 27 > 5
}


