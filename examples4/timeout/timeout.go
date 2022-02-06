package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "msg1"
	}()
	select {
	case value1 := <-c1:
		fmt.Println(value1)
	case <-time.After(time.Second * 1):
		fmt.Println("c1 timeout")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "msg2"
	}()
	select {
	case value2 := <-c2:
		fmt.Println(value2)
	case <-time.After(time.Second * 2):
		fmt.Println("c2 timeout")
	}
}
