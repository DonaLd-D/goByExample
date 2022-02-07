package main

import "fmt"
import "time"

func main() {
	ticker1 := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker1.C {
			fmt.Println(t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	ticker1.Stop()
	fmt.Println("ticker1 stopped")
}
