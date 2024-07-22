package main

import (
	"fmt"
	"time"
)

// infinite running go routine
func doWork(done <- chan bool) {
	for {
		select {
			case <- done:
				fmt.Println("done")
				return

			default:
				fmt.Println("running infinitely")
		}
	}
}

func main(){
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)
	close(done)
}
 