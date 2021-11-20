package main

import "fmt"

// func with parameter

func funcWithParameterDemo()  {
	fmt.Println(say("tirmizee"))
	fmt.Println(divise(5, 10))
	fmt.Println(add(5, 10, 30, 40))
}

func say(name string) string {
	return "Hello " + name
}

func divise(num1 int, num2 int) int {
	return num1 / num2
}

func add(num...int) int {
	result := 0
	for _, n  := range num {
		result += n
	}
	return result
}