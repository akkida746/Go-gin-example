package main
import ("fmt")

func main(){
	type Stack[T any] struct{
		vals []T
	}
	var s Stack[int]
	s.vals = append(s.vals, 1)
	s.vals = append(s.vals, 2)
	fmt.Println(s.vals)
}