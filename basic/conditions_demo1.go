package main

import "fmt"

func condisionsDemo1() {

	// &&
	// ||
	// !

	if 1 > 0 {
		fmt.Println(" 1 > 0 is true")
	}

	score := 99
	if score >= 80 {
		fmt.Println("Grade A")
	} else if score >= 70 {
		fmt.Println("Grade B")
	} else if score >= 60 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	x, y := 9, 9
	if x == y && x > 8 {
		fmt.Println("X = Y and X > 8")
	} else {
		fmt.Println(x)
	}

	z := 50
	if z > 0 || z <= 100 {
		fmt.Println("z between 0-100")
	}

}
