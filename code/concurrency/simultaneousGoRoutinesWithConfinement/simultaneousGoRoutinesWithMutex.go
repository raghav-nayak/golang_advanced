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

