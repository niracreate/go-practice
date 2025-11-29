package main

import (
	"fmt"
)

//A channel with a temporary storage (queue) inside it.
// How it behaves:

// Sender can send up to buffer size without blocking.

// Sender blocks only when buffer is full.

// Receiver blocks only when buffer is empty.

// sending
//
//	func processNum(numChan chan int) {
//		for num := range numChan {
//			fmt.Println("processing number" , num)
//			time.Sleep(time.Second)
//		}
//	}
//
// receiving
//
//	func sum(result chan int, num1 int, num2 int) {
//		numresult := num1 + num2
//		result <- numresult
//	}
//
// go routine syncronizer
// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("process")
// }
//
//
//<-done receives from the channel instead of sending.

//The goroutine will now wait until someone sends a value on done.

//chan<- int → send-only
//<-chan int → receive-only

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()

// 	for email := range emailChan {
// 		fmt.Println("email sending to", email)
// 		time.Sleep(time.Second)

// 	}

// }

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "test sending"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Value := <-chan1:
			fmt.Println("received data from chan 1", chan1Value)
		case chan2Value := <-chan2:
			fmt.Println("received data from chan 2", chan2Value)
		}
	}

	// emailChan := make(chan string, 5)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done setting")
	// close(emailChan)
	// <-done

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)
	// go task(done)
	// <-done
	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result
	// fmt.Println(res)
	// numChan := make(chan int)
	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// mesaageChan := make(chan string)   //blocking
	// mesaageChan <- "ping"
	// msg := <-mesaageChan
	// fmt.Println(msg)
}
