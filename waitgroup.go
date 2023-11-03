package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create buffered channels with a capacity of 1.
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)

	// Create a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// First goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch1 <- "Message from Goroutine 1"
	}()

	// Second goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch2 <- "Message from Goroutine 2"
	}()

	// Third goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(4 * time.Second)
		ch3 <- "Message from Goroutine 3"
	}()

	fmt.Println("blocking on waitgroup")

	// Wait for all goroutines to complete.
	wg.Wait()

	fmt.Println("unblocked")

	// Receive and print messages from all channels.
	msg1 := <-ch1
	msg2 := <-ch2
	msg3 := <-ch3

	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
}
