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

