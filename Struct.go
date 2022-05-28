package main

import "fmt"

func main() {
	var person struct{
		name string
		age int
	}
	person.name = "Adam"
	person.age = 20

	fmt.Println(person)
}