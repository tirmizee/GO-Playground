package main

import "fmt"

// func with return multiple value

func keyPair() (string, string) {
	return "key", "value"
}

func keyPair2(k string, v string) (string, string) {
	return k, v
}

func trpple(a int, b string, c bool) (int, string, bool) {
	return a, b, c
}

func returnMultipleDemo() {

	k, _ := keyPair()
	fmt.Println(k)

	_, v := keyPair()
	fmt.Println(v)

	key, value := keyPair()
	fmt.Println(key, value)

	key2, value2 := keyPair2("Hello", "World")
	fmt.Println(key2, value2)

	a, b, c := trpple(27, "tir", true)
	fmt.Println(a, b, c)

}
