package main

import "fmt"
import "time"

func worker(id int, from <-chan int, to chan<- int) {
	for f := range from {
		fmt.Println("worker", id, "processing from", f)
		time.Sleep(time.Millisecond * 500)
		to <- f * 2
	}

}

func main() {
	from := make(chan int, 50)
	to := make(chan int, 50)

	for i := 0; i < 3; i++ {
		go worker(i, from, to)
	}

	for j := 0; j < 9; j++ {
		from <- j
	}
	close(from)

	for k := 0; k < 9; k++ {
		<-to
	}
}
