https://www.youtube.com/watch?v=qyM8Pi1KiiM&list=PL7g1jYj15RUNqJStuwE9SCmeOKpgxC0HP

Concurrency can be achieved using 
1. go routines
2. channels
3. select

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
$ go run goRoutines.go
3
2
1
hi
```

note: the order execution is any order. 


Go uses ==fork-join model==
![[go_routine.png]]
