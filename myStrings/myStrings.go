package main

import (
	"fmt"
	"strings"
)

func main() {
	/*s := "Go string example"
	for k, v := range s {
		fmt.Printf("k：%d,v：%c == %d\n", k, &v, v)
	}*/

	s2 := "hel" + "lo, "
	s2 += "world!"
	fmt.Println(s2, "\n")

	s3 := strings.Join([]string{"hello", "world"}, ", ")
	fmt.Println(s3)
}
