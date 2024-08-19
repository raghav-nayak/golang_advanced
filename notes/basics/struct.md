struct
- group of related data
- collection of fields


```go
package main

import (
	"encoding/json"
	"fmt"
)

type Employee  struct {
	name string 
	age int
	isRemote bool
	address Address
}

func (e Employee) updateName(newName string)  {
	e.name = newName
	fmt.Println(e.name) // Jane
}

func (e *Employee) updateNameWithPtr(newName string)  {
	e.name = newName
}


type Address struct {
	Street string  `json:"street"`,
	City string `json:"city"`
}

func (a Address) printAddress() {
	fmt.Println("Street: ",a.Street, " City: ", a.City)
}


type Department struct {
	Name string `json:"name"`
	Employees []Employee `json:"emps"`
}

func (d Department) printDepartment() {
	fmt.Println("Name: ", d.Name)
	fmt.Println("Employees: ", d.Employees)
}


func main() {
	address := Address{
		Street:  "123 Main St",
		City:    "Maretown",
	}

	employee1 := Employee{
		name: "John",
		age: 30,
		isRemote: true,
		address: address,
	}

	fmt.Println(employee1) // {John 30 true}

	fmt.Println(employee1.address) // {123 Main St Maretown}
	fmt.Println(employee1.address.Street) // 23 Main St
	employee1.address.printAddress() // Street:  123 Main St  City:  Maretown

	// anonymous struct
	job := struct {
		title string
		salary int
	}{
		title: "Software Engineer",
		salary: 100000,
	}

	fmt.Println(job) // {Software Engineer 100000}

	employeePtr := &employee1
	employeePtr.age = 35
	fmt.Println(employee1) // {John 35 true}

	employee1.updateName("Jane")
	fmt.Println(employee1) // {John 35 true}

	employee1.updateNameWithPtr("Jack")
	fmt.Println(employee1) // {Jack 35 true}

	employee2 := Employee{
		name: "Johnson",
		age: 40,
		isRemote: false,
		address: Address {
			Street: "456 Main St",
			City: "Bronx",
		},
	}

	department := Department{
		Name: "IT",
		Employees: []Employee{
			employee1,
			employee2,
		},
	}

	deptJsonData, _ := json.Marshal(department)

	fmt.Println(deptJsonData) // [123 34 110 97 109 101 34 58 34 73 84 34 44 34 101 109 112 115 34 58 91 123 125 44 123 125 93 125]
	fmt.Println(string(deptJsonData)) {"name":"IT","emps":[{},{}]}
}
```



### struct tags 
- attach metadata to struct fields
- allows developers to use **custom annotations** for several use cases.
	- e.g. 
		- database mapping
		- validation
		- serialization
- allows us to use declarative programming pattern

reflection
- examine, introspect, and modify behavior **at runtime**
- disadvantages
	- less performant
	- hard to read
- use it when it is necessary

```go
type User struct {
	Name string `validate: "min=2,max=32"`
	Email string `validate: "required,email"`
}
```
you can give any key. Here we are using user defined key called `validate`. This is called tags.

```go
t := reflect.TypeOf(user)
fmt.Println("Name: ", t.Name()) // Name:  User
fmt.Println("Kind: ", t.Kind()) // Kind:  struct

for i:= 0; i < t.NumField(); i++ {
	field := t.Field(i)

	fmt.Println("field: ", field) // {Name  string validate: "min=2,max=32" 0 [0] false}
	fmt.Println("Name: ", field.Name) // Name:  Name
	fmt.Println("Tag: ", field.Tag) // Tag:  validate: "min=2,max=32"
}
```


