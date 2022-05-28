package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	var a [2]int
	a[0] = 1;
	fmt.Println(a)

	var b = [2]int{1,2}
	fmt.Println(b)
}
