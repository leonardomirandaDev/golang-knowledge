package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName string
	age int
}

// p is a copy
func (p Person) completeName() string {
	return p.firstName + " " + p.lastName
}


// p is exact object
func (p *Person) uppercaseName() {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}


func main() {
	var leonardo = Person{firstName: "Leonardo", lastName:"Miranda", age: 23}

	fmt.Println(leonardo.completeName())
	fmt.Println("-------")
	leonardo.uppercaseName();
	fmt.Println(leonardo.completeName())
}