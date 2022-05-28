package main

import "fmt"

func main(){
	fmt.Println("Functions")
	checkType(2)
	fmt.Println(getMod(3))
	fmt.Println("Passing function as parameter")
	calc(print)
}
func checkType(n int){
	if n%2==0 {
		fmt.Println("Even")
	} else{
		fmt.Println("Odd");
	}
}
func getMod(n int)int{
	return n%2;
}
func calc(f func()){
	f()
}
func print(){
	fmt.Println("Function called")
}
func printVal(n int){
	fmt.Println(n)
}