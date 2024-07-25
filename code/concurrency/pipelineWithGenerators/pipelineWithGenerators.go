package main

import (
	"fmt"
	"math/rand"
)

func repeatFunc[T any, K any](done <- chan K, fn func() T ) <- chan T {
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

func take[T any, K any](done <- chan K, stream <- chan T, num int) <- chan T {
	taken := make(chan T)

	go func() {
		defer close(taken)
		for i := 0; i < num; i++ {
			select {
				case <- done:
					return
				case taken <- <- stream:
			}
		}
	}()
	return taken
}



func main() {
	done := make(chan int)
	defer close(done)

	randomNumFetcher := func() int {
		return rand.Intn(5000000)
	}

	for randomNum := range take(done, repeatFunc(done, randomNumFetcher), 10){
		fmt.Println(randomNum)
	}
}