package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3) // we will start 3 goroutines

	// start 3 goroutines using anonymous functions
	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done() // mark this goroutine done when finished
			fmt.Printf("Goroutine %d is running\n", id)
		}(i) // pass i as parameter to avoid closure issue
	}

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines finished")
}
