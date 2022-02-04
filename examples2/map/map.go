package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("get map: ", v1)

	fmt.Println("length: ", len(m))

	delete(m, "k1")
	fmt.Println("get map: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("get map n: ", n)
}
