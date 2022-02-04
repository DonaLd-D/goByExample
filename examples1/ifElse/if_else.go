package main

import "fmt"

func main() {
	if 5%2 == 0 {
		fmt.Println("5 is even")
	} else {
		fmt.Println("5 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
