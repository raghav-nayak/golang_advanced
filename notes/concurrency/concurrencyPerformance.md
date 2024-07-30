#go #concurrency 

**critical section**
- there is no synchronous access between go routines.
- result in race condition
![[go_critical_section.png]]


Solution: Mutex(mutual exclusion) -> naive approach

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

##### confinement

