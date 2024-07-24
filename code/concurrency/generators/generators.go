package main

import (
	"fmt"
	"math/rand"
)

func repeatFunc[T any, K any](done <- chan K, fn func() T ) <- chan T {
	/*
	Function Parameters:

	done: This is a channel of type <-chan K. It receives values of any type K (often used for signaling).
	fn: This is a function that takes no arguments and returns a value of any type T.
	Function Return Value:

	The function returns a channel of type <-chan T. This channel will be used to receive the return values of the function fn.

	*/

	stream := make(chan T)
	go func() {
		defer close(stream) // ensures that the stream channel is closed after the goroutine finishes. This prevents further writes to the channel.
		for { // infinite loop
			select {
				case <- done: //  to stop execution
					return
				case stream <- fn(): // passing function to channel
			}
		}
	}()

	return stream
}

func main() {
	done := make(chan int)
	defer close(done)

	randomNumFetcher := func() int {
		return rand.Intn(5000000)
	}

	for randomNum := range repeatFunc(done, randomNumFetcher) {
		fmt.Println(randomNum)
	}
}