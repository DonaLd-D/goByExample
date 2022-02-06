package main

import "fmt"

func main() {
	msg := make(chan int, 2)
	msg <- 1
	msg <- 2
	close(msg)
	for m := range msg {
		fmt.Println("msg is ", m)
	}
}
