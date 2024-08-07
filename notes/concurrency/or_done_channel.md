

![[go_do_channel.png]]

![[go_or_done_channel.png]]



```go
package main

import "fmt"

func main() {
	done := make(chan interface{})

	done <- "string"
	done <- 24

	var1 := <- done
	var2 := <- done

	fmt.Println(var1)
	fmt.Println(var2)
}
```

output
```shell
$ go run or_done_channel.go 
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()

/go/src/github.com/golang_advanced/code/concurrency/or_done_channel/or_done_channel.go:8 +0x50
exit status 2
```

solution is 
`done := make(chan interface{}, 10)`

output
```shell
$ go run or_done_channel.go
string
24
```


```go
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
```
