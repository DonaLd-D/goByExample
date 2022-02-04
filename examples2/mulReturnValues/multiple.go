package main

import "fmt"

func add(x, y int) (x1, y2, res int) {
	return x, y, x + y
}

func main() {
	x, y, res := add(2, 3)
	fmt.Println(x, y, res)

	_, _, res1 := add(3, 4)
	fmt.Println(res1)
}
