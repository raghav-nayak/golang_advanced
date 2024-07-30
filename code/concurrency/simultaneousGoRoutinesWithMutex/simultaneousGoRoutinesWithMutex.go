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

