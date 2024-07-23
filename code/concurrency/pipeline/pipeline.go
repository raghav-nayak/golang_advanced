package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int) // unbuffered channel
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}() //invoke as it is a anonymous function
	return out
}

func square(in <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		for n:= range in { // blocked till it reads the value from the channel
			out <- n * n 
		}
	}()
	return out
}
func main() {
	// input
	nums := []int{2,34,7,1}

	// stage 1
	dataChannel := sliceToChannel(nums)

	// stage 2
	squaredChannel := square(dataChannel)

	// stage 3
	for n := range squaredChannel{
		fmt.Println(n)
	}
}
