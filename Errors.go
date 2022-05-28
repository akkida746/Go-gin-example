package main

import (
	"errors";"fmt"
)

func main(){
	fmt.Println("Errors")
	a, err := calc(0)
	fmt.Println(a, err)
	b, err := calc(1)
	fmt.Println(b, err)
}
func calc(n int)(int,error){
	if n==0{
		return 0, errors.New("number is zero")
	}
	if n==1{
		return 1, &CustomError{msg: "CustomMsg"}
	}
	return n, nil
}
//Custom error
type CustomError struct{
	msg string
}
func (e *CustomError) Error() string{
	return e.msg
}
