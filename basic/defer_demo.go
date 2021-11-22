package main

import "fmt"

func deferDemo() {

	fmt.Println("Hello")
	defer closeDatabase1()
	defer closeDatabase2()
	defer closeDatabase3()
	fmt.Println("World")

}

func closeDatabase1() {
	fmt.Println("close database 1")
}

func closeDatabase2() {
	fmt.Println("close database 2")
}

func closeDatabase3() {
	fmt.Println("close database 3")
}
