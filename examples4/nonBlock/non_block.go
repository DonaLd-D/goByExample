package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case res := <-messages:
		fmt.Println("messages from ", res)
	default:
		fmt.Println("no messages")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("messages from ", messages)
	default:
		fmt.Println("no messages")
	}

	select {
	case res := <-messages:
		fmt.Println("messages from ", res)
	case sig := <-signals:
		fmt.Println("signals from ", sig)
	default:
		fmt.Println("no activity")
	}
}
