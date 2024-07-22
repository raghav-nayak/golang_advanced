package main

import (
	"fmt"
	"time"
)

func main(){
	charChannel := make(chan string, 3) //buffered channel
	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		select {
			case charChannel <- char:
		}
	}
	close(charChannel)

	for char := range charChannel {
		fmt.Println(char)
	}

	// infinite running go routine
	go func() {
		for {
			select {
			default:
				fmt.Println("running infinitely")
			}
		}
	}()

	time.Sleep(time.Second * 2)
}
 