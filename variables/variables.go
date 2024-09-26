package main

import (
	"fmt"
)

/*func main() {
	x := 1
	fmt.Println(x) // prints 1
	{
		fmt.Println(x) // prints 1
		x := 2
		fmt.Println(x) // prints 2
	}
	fmt.Println("qwe %d", x) // prints 1
}*/

func main() {
	x := 1
	fmt.Print("x is ", x, "\n")
	fmt.Printf("x is %d\n", x)
	fmt.Println("x is", x)
}
