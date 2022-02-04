package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 3 {
			fmt.Println("index is: ", i)
		}
	}
	fmt.Println("sum of sums: ", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, s := range "golang" {
		fmt.Println(i, s)
	}
}
