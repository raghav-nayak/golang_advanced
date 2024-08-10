package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	done := make(chan interface{}, 10)

	defer close(done)

	cows := make(chan interface{}, 10)
	pigs := make(chan interface{}, 10)

	go func() {
		for {
			select {
			case <- done:
				return
			case cows <- "moo":
			}
		}
	}()
	
	go func() {
		for {
			select {
			case <- done:
				return
			case pigs <- "oink":
			}
		}
	}()

	wg.Add(1)
	go consumeCows(done, cows)
	wg.Add(1)
	go consumePigs(done, pigs)

}

func consumeCows(done <-chan interface{}, cows <-chan interface{}){
	defer wg.Done()

	for {
		select {
		case <- done:
			return
		case cow, ok := <- cows:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			// do some complex logic
			fmt.Println(cow)
		}
	}
} 


func consumePigs(done <-chan interface{}, pigs <-chan interface{}){
	defer wg.Done()

	for {
		select {
		case <- done:
			return
		case pig, ok := <- pigs:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			// do some complex logic
			fmt.Println(pig)
		}
	}
} 

func orDone(done, ch <-chan interface{}) <-chan interface{} {
	relayStream := make(chan interface{}) // unbuffered channel

	go func() {
		defer close(relayStream)
		for {
			select {
			case <- done:
				return
			case value, ok := <- ch:
				if !ok {
					return
				}
				select { // this is important to prevent blocking
				case relayStream <- value:
				}
			}

		}
	}()
	return relayStream
}