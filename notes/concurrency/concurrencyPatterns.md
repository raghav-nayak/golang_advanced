### 1. for-select loop

check [[buffered vs. unbuffered channels]]

```go
package main

import "fmt"

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
}
```


```go
package main

import (
	"fmt"
	"time"
)

func main(){
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
```


### 2. Done channel
the parent or caller function can control the infinite running go routine by using done channel.

```go
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
	close(done) // this will initial the stopping of infinitely running go routine
}
```


### 3. Pipeline

![[go_pipeline.png]]

We can achieve something in each stages. We separate the concern in each stages.

```go
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
```

pipeline with generators

![[go_pipeline_with_generators_1.png]]


![[go_pipeline_with_generators_2.png]]


```go
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
```

Now you will get 10 random integer numbers only.
