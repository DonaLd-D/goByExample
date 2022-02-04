package main

import "fmt"

func add(nums ...int) {
	fmt.Println("nums is: ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total is: ", total)
}

func main() {
	add(2, 3, 4)

	add(4, 5, 6, 7)

	nums := []int{12, 34, 56}
	add(nums...)
}
