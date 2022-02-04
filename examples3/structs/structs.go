package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	s := person{"jack ma", 60}
	fmt.Println(s)
}
