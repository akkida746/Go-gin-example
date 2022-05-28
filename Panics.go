package main

import "fmt"

func main(){
	fmt.Println("Panics")
	a := []int{1,2}
	fmt.Println(a[1])
	//fmt.Println(a[2])// This will return panic: runtime error
	callPanic()
}
func callPanic(){
	//this method will be called even after panic is occured
	// defer is used to run code even when panic is occured, this is similar to finally in java
	defer func(){
		fmt.Println("Handled Panic")
	}()
	divide(1,0)
	fmt.Println("Completed")
}
func divide(a, b int)int{
	return a/b
}