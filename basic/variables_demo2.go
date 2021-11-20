package main

import "fmt"

func declareMultiVariableDemo() {

	var text, text2 string
	fmt.Println(text, text2)

	var length, size = 10, 10
	fmt.Println(length, size)

	isTrue, amount := true, 209.11
	fmt.Println(isTrue, amount)

	first, second, third := "name", "mid", "last"
	fmt.Println(first, second, third)

	_, _, _, last := 0, 1, 2, 3
	fmt.Println(last)

}
