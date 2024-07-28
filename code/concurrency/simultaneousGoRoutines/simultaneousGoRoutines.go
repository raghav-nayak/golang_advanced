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

