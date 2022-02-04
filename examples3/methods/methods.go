package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{10, 4}
	fmt.Println("area is: ", r.area())
	fmt.Println("perimeter is: ", r.perimeter())
}
