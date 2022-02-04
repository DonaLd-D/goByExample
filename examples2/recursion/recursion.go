package main

import "fmt"

func add(x int) int {
	if x <= 1 {
		return 1
	}
	return x * add(x-1)
}

func main() {
	res := add(7)
	fmt.Println(res)
}
