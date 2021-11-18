package main

import "fmt"

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

	numbers7 := numbers[:]
	fmt.Println(numbers7)

}