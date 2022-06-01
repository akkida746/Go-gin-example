package main

import ("fmt";"time")

func main(){
	ch := make(chan int)
	go func(){
		for i:=1;i<5;i++{
			ch<-i
			time.Sleep(time.Millisecond*1000)
		}
		close(ch)
	}()

	for i:=1;i<5;i++{
		fmt.Println(<-ch) // Blocking call
	}
}