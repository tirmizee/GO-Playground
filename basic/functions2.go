package main

import "fmt"

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