package main

import (
	"fmt"
	"time"
)

func main() {
	msg1 := make(chan string)
	msg2 := make(chan string)

	go func(s string) {
		time.Sleep(time.Second)
		msg1 <- s
	}("msg1")

	go func(s string) {
		time.Sleep(time.Second * 1)
		msg2 <- s
	}("msg2")

	for i := 0; i < 2; i++ {
		select {
		case c1 := <-msg1:
			fmt.Println("received:", c1)
		case c2 := <-msg2:
			fmt.Println("received:", c2)
		}
	}
}
