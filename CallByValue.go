package main

import "fmt"

func main(){
	fmt.Println("CallByValue")
	s := "hello"
	fmt.Println(s)
	change(s)
	fmt.Println(s)
}
func change(s string){
	s = "world"
	fmt.Println("Changed value in function: " + s)
}