package main

import "fmt"
import "errors"

func f1(num int) (int, error) {
	if num%2 == 1 {
		return -1, errors.New("fail")
	}
	return num, nil
}

type argError struct {
	arg  int
	prob string
}

func (arg argError) Error() string {
	return fmt.Sprintf("%d -> %s", arg.arg, arg.prob)
}

func f2(num int) (int, error) {
	if num%2 == 1 {
		return -1, &argError{num, "fail"}
	}
	return num, nil
}

func main() {
	for _, num := range []int{23, 34} {
		if i, r := f1(num); r != nil {
			fmt.Println(r)
		} else {
			fmt.Println(i)
		}
	}

	for _, num := range []int{23, 34} {
		if i, r := f2(num); r != nil {
			fmt.Println(r)
		} else {
			fmt.Println(i)
		}
	}
}
