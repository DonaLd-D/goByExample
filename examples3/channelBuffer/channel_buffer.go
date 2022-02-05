package main

import "fmt"

func main() {
	msg := make(chan string, 2)

	msg <- "channels"
	msg <- "buffer"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
