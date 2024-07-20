channels
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


