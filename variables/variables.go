package main

import (
	"fmt"
	"time"
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

/*func main() {
	x := 1
	fmt.Print("x is ", x, "\n")
	fmt.Printf("x is %d\n", x)
	fmt.Println("x is", x)
}*/

func main() {
	weights := []int{
		45, 14, 73, 18, 92, 68, 3, 10,
	}
	fmt.Printf("now is %d year\n", dateNow())
	fmt.Printf("the average weight of your friends %d year\n", meanWeight(weights))
}

func dateNow() int {
	d := time.Date(1994, time.April, 27, 0, 0, 0, 0, time.UTC)
	year, _, _ := d.Date()
	return year + 30
}

func meanWeight(weights []int) int {
	var total_weights, mean_weight int
	for _, weight := range weights {
		total_weights += weight
	}
	mean_weight = total_weights / len(weights)
	return mean_weight
}
