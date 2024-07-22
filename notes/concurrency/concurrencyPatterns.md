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
The **done channel** is commonly used to signal the completion of a goroutine or a task. It is a channel that is closed when the task is finished.

The parent or caller function can control the infinite running go routine by using done channel.

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
