package main
import ("fmt")

func main(){
	type Foo struct{
		name string
	}

	var x Foo
	x.name = "Akash"

	fmt.Println(x.name)
}