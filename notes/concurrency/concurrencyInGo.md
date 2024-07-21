#golang #concurrency
https://www.youtube.com/watch?v=qyM8Pi1KiiM&list=PL7g1jYj15RUNqJStuwE9SCmeOKpgxC0HP

Concurrency can be achieved using 
1. go routines
2. channels
3. select


### 1. Go routines

```go
package main

import "fmt"

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	someFunc("2")
	fmt.Println("hi")
}
```

when we call `someFunc()`, the main will wait for `someFunc()` to complete its execution. So, the execution is blocked in this case.

to use go routines, we just need to add `go` before the function call.

As go routines are spawned and run independently. 

```go
func main() {
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")
	time.Sleep(time.Second * 1)
	fmt.Println("hi")
}
```

output:
```shell
$ go run primitives.go
3
2
1
hi
```

note: the order execution is any order. 


Go uses ==fork-join model==
![[go_routine.png]]


### 2. Channels
- used to communicate between go routines
- a mechanism to communicate between two or more go routine as go routines run independent of each other
- works as FIFO queues
- even main function can communicate with go routines using channels

```go
package main

import "fmt"

func main() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "hello"
	}()

	msg := <-myChannel //blocking 

	fmt.Println(msg)
}
```

In this example, we see the both main and func() communicate using channel. That means both main and func() are in sync now as `msg := <-myChannel` is a blocking call.

### 3. select

using select statement, we can read messages from more than one channels.
It blocks execution till it reads a message from one of the channels.
If multiple messages are ready, it randomly reads one of them.

```go
package main

import "fmt"

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "hello"
	}()
	
	go func() {
		anotherChannel <- "hi"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}
```
