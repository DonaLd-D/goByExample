package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("received value is ", job)
			} else {
				fmt.Println("no more data")
				done <- true
				return
			}
		}
	}()

	for i := 1; i < 4; i++ {
		jobs <- i
		fmt.Println("sent value ", i)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
