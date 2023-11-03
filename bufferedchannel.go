package main

import "fmt"

func main() {
	// Create a buffered channel with a capacity of 2.
	ch := make(chan int, 2)

	// Send two values without blocking.
	ch <- 1
	ch <- 2

	// This third send will block since the channel buffer is full.
	// Uncomment this line to see the blocking behavior.
	ch <- 3
	// throws a 'fatal error: all goroutines are asleep - deadlock!'

	// Receive and print the values from the channel.
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Now, the channel has space, so the send won't block.
	ch <- 3

	// Receive and print the third value from the channel.
	fmt.Println(<-ch)
}
