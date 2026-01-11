package main

import (
	"fmt"
	"time"
)

func worker1(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "start...", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "end...", j)
		results <- j * 2
	}
}

func main() {
	const numjobs = 10
	jobs := make(chan int, numjobs)
	results := make(chan int, numjobs)
	for w := 1; w <= 3; w++ {
		go worker1(w, jobs, results)
	}
	for j := 1; j <= numjobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numjobs; a++ {
		<-results
	}
}
