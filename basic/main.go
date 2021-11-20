package main

import (
	"basic/another"
)

func main() {
	var _ another.Linklist
	var _ another.desk // error desk not exported by package another

	var _ another.Employee
	var _ another.account // error account not exported by package another

	another.RandomNumber()
	another.randonText() // error randonText not exported by package another
}
