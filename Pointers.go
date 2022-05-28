package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	a := 1
	fmt.Println(a)
	x := &a
	*x = 2
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(*x)
}
