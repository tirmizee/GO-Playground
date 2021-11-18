package main

import "fmt"

// arrays 

func arraysStringDemo()  {

	var students1 [3]string
	students1[0] = "students1 a"
	students1[1] = "students1 b"
	students1[2] = "students1 c"
	fmt.Println(students1[0])
	fmt.Println(students1[1])
	fmt.Println(students1[2])


	var students2 [3]string = [3]string {
		"students2 a",
		"students2 b",
		"students2 c",
	}
	fmt.Println(students2[0])
	fmt.Println(students2[1])
	fmt.Println(students2[2])

	var students3 []string = []string {
		"students3 a",
		"students3 b",
		"students3 c",
	}
	fmt.Println(students3[0])
	fmt.Println(students3[1])
	fmt.Println(students3[2])

	students4 := [3]string {
		"students4 a",
		"students4 b",
		"students4 c",
	}
	fmt.Println(students4[0])
	fmt.Println(students4[1])
	fmt.Println(students4[2])

	students5 := []string {
		"students5 a",
		"students5 b",
		"students5 c",
	}
	fmt.Println(students5[0])
	fmt.Println(students5[1])
	fmt.Println(students5[2])

	students6 := make([]string, 3)
	students6[0] = "students6 a"
	students6[1] = "students6 b"
	students6[2] = "students6 c"
	fmt.Println(students6)

}
