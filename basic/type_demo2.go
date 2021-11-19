package main

import "fmt"

// style 1
type Person struct{ 
	name, city, phone  string 
	age                int 
}

/*
// style 2
type Person struct {
	name string
	city string
	phone string
	age int
}
*/

func demoType2_1()  {
	person := Person{
		name: "tirmizee",
		age: 27,
		phone: "0999999999",
		city: "Bankok",
	}
	fmt.Println(person)
}

func demoType2_2()  {
	person := Person{
		"tirmizee",
		"Bankok",
		"0999999999",
		27,
	}
	fmt.Println(person)
}

func demoType2_3()  {
	// var person Person = Person{}
	// var person Person
	// person := Person{}
	person := Person{}
	person.name = "tirmizee"
	person.age = 27
	person.phone = "0999999999"
	person.city = "Bankok"
	fmt.Println(person)
}
