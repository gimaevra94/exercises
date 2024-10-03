package main

import (
	"fmt"
)

type Person struct{
	name string
	job string
}

func main() {
	var aperson Person

	aperson.name="хуй"
	aperson.job="пизда"

	fmt.Printf("\nимя - %s\nработа - %s\n\n", aperson.name,aperson.job)
}