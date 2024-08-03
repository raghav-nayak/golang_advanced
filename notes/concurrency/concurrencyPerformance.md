#go #concurrency 

**critical section**
- there is no synchronous access between go routines.
- result in race condition
![[go_critical_section.png]]


Solution: Mutex(mutual exclusion) -> naive approach

#### Mutex

A mutex (short for "mutual exclusion") in Go is a fundamental synchronization primitive used to protect shared resources from concurrent access. It ensures that only one goroutine can access a critical section of code at a time, thus preventing race conditions and ensuring data consistency.

##### Basic Usage

The `sync` package in Go provides two main types of mutexes:

1. **Mutex** (`sync.Mutex`)
2. **RWMutex** (`sync.RWMutex`)


A `Mutex` provides a simple locking mechanism.

- **Lock**: The `Lock` method locks the mutex. If the mutex is already locked, the calling goroutine blocks until the mutex is unlocked.
- **Unlock**: The `Unlock` method unlocks the mutex. If other goroutines are waiting for the mutex to be unlocked, one of them will proceed to lock it.


It is kind of locking mechanism which locks the usage of the resource and other go routines must wait till the lock is released.
It is a naive approach.
This will create a bottleneck.


without any locking mechanism

`simultaneousGoRoutines.go`
```go
package main

import (
	"fmt"
	"sync"
)

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done() // to inform the go routine that it is done
	processedData := data * 2
	*result = append(*result, processedData)
}

func main() {
	var wg sync.WaitGroup

	input := []int{1,2,3,4,5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data) // wg is passed to tell the function that it is finished
	}

	wg.Wait()
	fmt.Println(result)
}
```

output
```sh
$ go run simultaneousGoRoutines.go               
[10 6 8 2]

$ go run simultaneousGoRoutines.go 
[2 6 4]
```


with mutex(naive solution)
```Go
package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	lock.Lock()
	defer wg.Done() // to inform the go routine that it is done
	processedData := data * 2
	*result = append(*result, processedData)
	lock.Unlock()
}


func main() {
	var wg sync.WaitGroup

	input := []int{1,2,3,4,5}

	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data) // wg is passed to tell the function that it is finished

	}
	wg.Wait()
	fmt.Println(result)
}
```

output:
```shell
$ go run simultaneousGoRoutinesWithMutex.go
[10 6 8 4 2]

$ go run simultaneousGoRoutinesWithMutex.go
[10 6 8 4 2]

$ go run simultaneousGoRoutinesWithMutex.go
[10 4 6 8 2]

$ go run simultaneousGoRoutinesWithMutex.go
[10 2 4 6 8]

$ go run simultaneousGoRoutinesWithMutex.go
[4 10 2 6 8]

$ go run simultaneousGoRoutinesWithMutex.go
[2 10 4 6 8]
```
As you can see, we get all the numbers 


If there is any processing between lock and unlock, it takes time to process.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func process(data int) int {
	time.Sleep(time.Second * 2)
	return data * 2
}

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	lock.Lock()
	defer wg.Done() // to inform the go routine that it is done
	processedData := process(data)
	*result = append(*result, processedData)
	lock.Unlock()
}


func main() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1,2,3,4,5}

	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data) // wg is passed to tell the function that it is finished

	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(start))
}
```

==without lock==
```shell
$ go run simultaneousGoRoutinesWithMutex.go
[6 4]
2.001703625s
```

==with lock==
```sh
$ go run simultaneousGoRoutinesWithMutex.go
[2 4 10 8 6]
10.004206584sz
```

as you can see, it is taking 10 sec as each go routine take 2 sec to release the lock. Now, it is synchronous and iterative method even though we are using concurrency.

##### improved mutex

If you carefully observe, `processedData := process(data)` is not a critical section here. The critical section is `*result = append(*result, processedData)`.
**be careful about the locking any code**

By moving the lock, we can achieve better results.

```go
...
func processData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done() // to inform the go routine that it is done
	processedData := process(data)
	lock.Lock()
	*result = append(*result, processedData) // critical section
	lock.Unlock()
}
...
```

output
```shell
$ go run simultaneousGoRoutinesWithMutex.go
[4 2 10 6 8]
2.001387375s
```



#### Confinement 
confine the goroutine to specific part of the shared resource
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func process(data int) int {
	time.Sleep(time.Second * 2)
	return data * 2
}

func processData(wg *sync.WaitGroup, resultDest *int, data int) {
	defer wg.Done() // to inform the go routine that it is done
	processedData := process(data)

	*resultDest = processedData //critical section
}


func main() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1,2,3,4,5}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data)  

	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println(time.Since(start))
}
```

```shell
$ go run simultaneousGoRoutinesWithMutex.go 
[2 4 6 8 10]
2.000412459s

$ go run simultaneousGoRoutinesWithMutex.go
[2 4 6 8 10]
2.000947375s
```

Now you can see instead of passing entire slice, we are passing only one element. Now, you can see the output and order is maintained.

You can also check the race condition
```sh
$ go run -race simultaneousGoRoutinesWithMutex.go
[2 4 6 8 10]
2.001568875s
```
