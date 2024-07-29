#go #concurrency 

critical section
- there is no synchronous access between go routines.
- result in race condition
![[go_critical_section.png]]


Solution: Mutex(mutual exclusion)

It is kind of locking mechanism which locks the usage of the resource and other go routines must wait till the lock is released.


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
