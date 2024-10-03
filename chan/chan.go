package main

import (
    "fmt"
)

func producer(c chan int) {
    for i := 0; i < 5; i++ {
        c <- i
    }
    close(c)
}

func main() {
    c := make(chan int, 2)
    go producer(c)

    for i := range c {
        fmt.Println(i)
    }
}