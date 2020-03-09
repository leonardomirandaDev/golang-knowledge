package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	age int
}

func (p Person) completeName() string {
	return p.firstName + " " + p.lastName
}
func main() {
	var leonardo = Person{firstName: "Leonardo", lastName:"Miranda", age: 23}

	fmt.Println(leonardo.completeName())
}