package main

import (
	"fmt"
	"math/rand"
	"time"
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


func primeFinder(done <- chan int, randomIntStream <- chan int) <- chan int {

	isPrime := func(randomInt int) bool {
		for i:= 2; i < randomInt; i++ {
			if randomInt % i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
				case <- done:
					return
				case randomInt := <- randomIntStream:
					if isPrime(randomInt) {
						primes <- randomInt
					}
			}
		}
	}()
	return primes
}


func main() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	randomNumFetcher := func() int {
		return rand.Intn(10000000)
	}


	randomIntStream := repeatFunc(done, randomNumFetcher)

	primeStream := primeFinder(done, randomIntStream)

	for randomNum := range take(done, primeStream, 10){
		fmt.Println(randomNum)
	}

	fmt.Println(time.Since(start))

}