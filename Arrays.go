package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	var a [2]int
	a[0] = 1
	fmt.Println(a)

	var b = [2]int{1, 2} // constant array and non-growable
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(b)

	var c []int // growable array as we have not defined size here.
	c = append(c, 1)
	c = append(c, 2)
	c = append(c, 3)
	c = append(c, 4)
	fmt.Println(c)
	fmt.Println("Printing slicing")
	fmt.Println(c[:2])
	fmt.Println(c[:1])
	fmt.Println(c[1:])
}
