package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a buffered channel with a capacity of 1.
	ch := make(chan string, 1)

	// First goroutine
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Message from Goroutine 1"
	}()

	// Second goroutine
	go func() {
		time.Sleep(4 * time.Second)
		ch <- "Message from Goroutine 2"
	}()

	// Third goroutine
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Message from Goroutine 3"
	}()

	fmt.Println("blocked")

	// Select statement to receive from the first sender and unblock
	select {
	case msg := <-ch:
		fmt.Println(msg)
	}

	fmt.Println("unblocked")

	// Close the channel to unblock other goroutines.
	close(ch)
}
