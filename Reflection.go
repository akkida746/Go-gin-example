package main

import (
	"fmt"
	"reflect"
)

func main(){
	type Foo struct{
		name string
	}

	var x Foo
	x.name = "Akash"
	fmt.Println(x.name)
	fmt.Println(reflect.TypeOf(x))
}