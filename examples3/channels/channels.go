package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "from channels"
	}()

	msg := <-ch
	fmt.Println(ch)
	fmt.Println(msg)
}
