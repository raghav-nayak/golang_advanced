# "or done" pattern
In Go, channels are often used for communication between goroutines, and the "or done" pattern helps manage the scenario where you need to wait for either a task to complete or for a cancellation signal to occur. The `select` statement is used to listen on multiple channels simultaneously, and this pattern leverages that capability.

When you have a channel for a task and a "done" channel for cancellation or completion signals, you can use the `select` statement to handle either case.

#### Use Cases
- **Handling Task Completion**: When you need to perform a task and wait for it to complete while also being able to handle timeouts or cancellation.
- **Graceful Shutdown**: When you need to stop ongoing operations when a shutdown signal is received, such as when handling server requests or background jobs.

### Key Points

- The "or done" pattern allows you to handle multiple conditions concurrently, making your code more responsive and flexible.
- Using channels for cancellation and task signaling is a common idiom in Go, leveraging the concurrency model effectively.
- Always be mindful of closing channels and properly handling all possible cases to avoid leaks or deadlocks.

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


#### goroutine leak
A goroutine leak occurs when a goroutine in a Go program is no longer needed but continues to run, consuming resources and potentially causing performance issues or crashes. Goroutines are lightweight threads managed by the Go runtime, and while they're designed to be efficient, they can still cause problems if not managed properly.

### Causes of Goroutine Leaks:

1. **Blocking Operations**: If a goroutine is blocked indefinitely on a channel or a synchronization primitive (like a mutex), it can lead to a leak if the goroutine is not supposed to be running forever.
2. **Uncleared Channels**: Goroutines waiting on channels that are never closed or receive values can lead to leaks, especially if the channel is not properly cleaned up when it's no longer needed.
3. **Improper Synchronization**: Using synchronization mechanisms incorrectly, such as not using `sync.WaitGroup` properly, can leave goroutines running after they should have been terminated.
4. **Unbounded Resource Usage**: Goroutines that continuously allocate memory or resources without proper limits or cleanup can lead to leaks.
    

### Detecting Goroutine Leaks:

1. **Profiling**: Go's built-in profiling tools, like `pprof`, can help identify goroutines that are stuck or not terminating. You can use these tools to generate profiles and analyze them to spot leaks.
2. **Static Analysis**: Tools like `golangci-lint` can sometimes detect potential issues in code that might lead to goroutine leaks.
3. **Code Review**: Carefully reviewing code for proper goroutine management practices can help catch potential leaks. Ensuring that channels are closed, synchronization is handled correctly, and goroutines are only running as long as needed is key.
### Mitigating Goroutine Leaks:

1. **Proper Channel Management**: Ensure that channels are closed properly when they are no longer needed.
2. **Use Contexts**: Using `context.Context` for managing goroutine lifetimes can help ensure they are canceled or terminated appropriately.
3. **Synchronization**: Use synchronization primitives like `sync.WaitGroup` to ensure all goroutines complete before the program exits.
4. **Timeouts and Limits**: Implement timeouts and resource limits to avoid goroutines running indefinitely.
5. **Testing**: Regularly test for potential leaks using stress tests and profiling tools.
    

By being mindful of these practices, you can minimize the risk of goroutine leaks and ensure your Go applications run efficiently.

in the above example, we are using `case <- done` to close the channel gracefully.

