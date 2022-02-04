package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	res := add(1, 2)
	fmt.Println("res is: ", res)
}
