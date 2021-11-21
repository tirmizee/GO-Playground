package main

import "fmt"

func main() {
	names := []string{}
	fmt.Printf("Length %d, Capacity %d \n", len(names), cap(names))

	names = append(names, "java")
	fmt.Printf("Length %d, Capacity %d \n", len(names), cap(names))

	names = append(names, "Kotlin")
	fmt.Printf("Length %d, Capacity %d \n", len(names), cap(names))

	names = append(names, "Golang")
	fmt.Printf("Length %d, Capacity %d \n", len(names), cap(names))
}
