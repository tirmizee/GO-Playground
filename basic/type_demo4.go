package main

import "fmt"

type Address struct {
	province    string
	district    string
}

type Customer struct {
	firstName   string
	lastName    string
	address     Address
}

func demoType4_1() {
	address := Address{
		province: "Bankok",
		district: "Maung",
	}

	customer := Customer{
		"Pratya",
		"Yeekhaday",
		address,
	}
	fmt.Println(customer)
}

func demoType4_2() {
	customer := Customer{
		"Pratya",
		"Yeekhaday",
		Address{
			province: "Bankok",
			district: "Maung",
		},
	}
	fmt.Println(customer)
}

func demoType4_3() {

	address := Address{}
	address.province = "Bankok"
	address.district = "Maung"

	var customer Customer = Customer{}
	customer.firstName = "Tirmizee"
	customer.lastName = "Gee"
	customer.address = address
	fmt.Println(customer)

}