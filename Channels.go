package main

import ("fmt";"time";"sync")

func main(){
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	go func(){
		for i:=1;i<5;i++{
			ch <- i
			time.Sleep(time.Millisecond * 1000)
		}	
		wg.Done()
	}()

	go func(){
		for i:=1;i<5;i++{
			a := <- ch
			fmt.Println(a)
		}	
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Completed");
}