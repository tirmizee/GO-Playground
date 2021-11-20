package main

import "fmt"

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func mapDemo1()  {

	map1 := map[string]string {
		"C01": "Red",
		"C02": "Yellow",
		"C03": "Blue",
	}
	fmt.Println(map1["C01"])

	map2 := make(map[string]string)
	map2["C01"] = "Red"
	map2["C02"] = "Yellow"
	map2["C03"] = "Blue"
	fmt.Println(map2["C02"])

	var map3 map[string]string
	map3 = map[string]string{}
	map3["C01"] = "Red"
	map3["C02"] = "Yellow"
	map3["C03"] = "Blue"
	fmt.Println(map3["C03"])

	printMap(map1)

}

