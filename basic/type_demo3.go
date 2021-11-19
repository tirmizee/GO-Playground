package main

import "fmt"

func demoType3_1()  {
	person := struct {
		name string
		city string
		phone string
		age int
	}{
		name: "tirmizee",
		age: 27,
		phone: "0999999999",
		city: "Bankok",
	}
	fmt.Println(person)
}