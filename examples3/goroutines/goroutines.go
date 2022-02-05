package main

import "fmt"

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, " ", i)
	}
}

func main() {
	f("start ")

	go f("start1 ")

	go func(s string) {
		fmt.Println(s)
	}("start3")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
