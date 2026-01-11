package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(ind int) {
	fmt.Printf("Worker %d starting...\n", ind)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done.\n", ind)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker2(i)
		}()
	}
	wg.Wait()
}
