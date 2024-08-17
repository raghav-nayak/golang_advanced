package main

import (
	"fmt"
	"math"
)

func main() {
	// variable
	var myVariable int = 10;
	myNewVariable := 10

	fmt.Println("myVariable: ", myVariable)
	fmt.Println("myNewVariable: ", myNewVariable)

	// array
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("myArray: ", myArray)

	// slice
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("mySlice: ", mySlice)

	mySlice = append(mySlice, 6)
	fmt.Println("appended mySlice: ", mySlice)

	// map
	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("myMap: ", myMap)

	// map using make
	myMakeMap := make(map[string]int)
	myMakeMap["a"] = 1
	myMakeMap["b"] = 2
	myMakeMap["c"] = 3
	fmt.Println("myMakeMap: ", myMakeMap)

	delete(myMap, "a")
	fmt.Println("after deleted, myMap: ", myMap)

	// empty make
	emptyMap := make(map[string]int)
	fmt.Println("emptyMap: ", emptyMap)

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	// for {

	// }

	// supports break and continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	sqrtResult, err := sqrt(1)
	if err != nil {
		panic("")
	}
	fmt.Println(sqrtResult)

	p := person{"John", 30}
	fmt.Println(p) // {John 30}
	fmt.Println(p.name) // John
	printStringer(p) // person struct
	p.setName("Jane") 
	fmt.Println(p) // {John 30}
	p.setNameWithPointer("Jack")
	fmt.Println(p) // {Jack 30}


	// pointer
	myVar := 0
	myPointer := &myVar
	fmt.Println("myVar = ", myVar, " myPointer = ", myPointer)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("a cannot be less than 0")
	}
	return math.Sqrt(a), nil
}

// struct
type person struct {
	name string
	age int
}


func printStringer(s stringer) {
	fmt.Println(s.string())
}

// interface
type stringer interface {
	string() string
}

func (p person) string() string {
	return "person struct"
}

func (p person) setName(newName string) {
	p.name = newName
}

func (p *person) setNameWithPointer(newName string) {
	p.name = newName
}