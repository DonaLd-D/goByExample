package main

import "fmt"

func zeroVal(x int) {
	x = 0
}

func zeroPtr(x *int) {
	*x = 0
}

func main() {
	i := 1
	fmt.Println("initial value is: ", i)
	zeroVal(i)
	fmt.Println("zero value is: ", i)

	zeroPtr(&i)
	fmt.Println("zero pointer is: ", i)

	fmt.Println("pointer is: ", &i)
}
