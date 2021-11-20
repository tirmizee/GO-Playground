package main

import "fmt"

func loopDemo() {

	names := []string{"Java", "C++", "GO", "Kotlin"}

	// for i
	for i := 0; i < len(names); i++ {
		fmt.Print(names[i], " ")
	}

	// foreach
	for i, n := range names {
		fmt.Print(i, n)
	}

	// while
	i := 0
	for i < len(names) {
		fmt.Print(names[i], " ")
		i++
	}

}
