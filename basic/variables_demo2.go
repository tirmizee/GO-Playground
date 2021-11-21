package main

import "fmt"

func declareMultiVariableDemo() {

	var text, text2 string
	fmt.Println(text, text2)

	var length, size = 10, 10
	fmt.Println(length, size)

	isTrue, amount := true, 209.11
	fmt.Println(isTrue, amount)

	_, second, third := "name", "mid", "last"
	fmt.Println(second, third)

	first, _, _, last := 0, 1, 2, 3
	fmt.Println(first, last)

}
