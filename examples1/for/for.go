package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		break
	}
}
