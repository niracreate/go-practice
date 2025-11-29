package main

import (
	"fmt"
	"time"
)

// A channel with no storage space inside.

// ch := make(chan int)   // unbuffered

// How it behaves:

// Sender waits until a receiver is ready.

// Receiver waits until a sender sends.

// They must meet at the same time → this is called blocking.


func main() {
	// Unbuffered channel (size 0)
	emailChan := make(chan string) // no buffer!

	// Start a worker goroutine that receives emails
	go func() {
		fmt.Println(<-emailChan) // receives first email
		fmt.Println(<-emailChan) // receives second email
		// If there were more, it would wait here
	}()

	// These sends now BLOCK until the worker receives them
	emailChan <- "1@example.com"
	emailChan <- "2@example.com"

	// Give the worker time to finish printing
	// (in real apps you use sync.WaitGroup instead)
	fmt.Println("All emails sent, waiting a bit...")
	time.Sleep(time.Second) // simple way for this demo
	fmt.Println("Main exits")
}
