package main

import "fmt"

// arrays 

func arraysStringDemo1()  {
	var students [3]string
	students[0] = "students1 a"
	students[1] = "students b"
	students[2] = "students c"
	fmt.Println(students[0])
	fmt.Println(students[1])
	fmt.Println(students[2])
}


func arraysStringDemo2() {
	var students [3]string = [3]string {
		"students a",
		"students b",
		"students c",
	}
	fmt.Println(students[0])
	fmt.Println(students[1])
	fmt.Println(students[2])
}

func arraysStringDemo3() {
	var students []string = []string {
		"students a",
		"students b",
		"students c",
	}
	fmt.Println(students[0])
	fmt.Println(students[1])
	fmt.Println(students[2])
}


func arraysStringDemo4() {
	students := [3]string {
		"students a",
		"students b",
		"students c",
	}
	fmt.Println(students[0])
	fmt.Println(students[1])
	fmt.Println(students[2])
}

func arraysStringDemo5() {
	students := []string {
		"students a",
		"students b",
		"students c",
	}
	fmt.Println(students[0])
	fmt.Println(students[1])
	fmt.Println(students[2])
}

func arraysStringDemo6() {
	students := make([]string, 3)
	students[0] = "students a"
	students[1] = "students b"
	students[2] = "students c"
	fmt.Println(students)
}