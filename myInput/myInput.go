package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("are you gay? ")
	answer, _ := reader.ReadString('\n')
	answer = "yes\n\n"
	fmt.Print(answer)

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("хуй длинный? ")
	answer2, _ := reader2.ReadString('\n')
	answer2 = "у тебя пизда"
	fmt.Print(answer2)
}
