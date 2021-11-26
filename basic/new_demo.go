package main

import (
	"fmt"
	"reflect"
)

type Vector struct {
	x int
	y int
}

func newDemo() {
	v := &Vector{}
	x := new(Vector)
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(reflect.TypeOf(x))
}
