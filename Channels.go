package main

import ("fmt";"time";"sync")

func main(){
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int)
	a := 1
	go func(){
		fmt.Println(a)
		for i:=0;i<3;i++{
			ch <- i
		}	
		wg.Done()
		close(ch)
	}()

	go func(){
		// while loop
		for true{
			a,isOpen := <-ch // Taking data from channel and checking if channel is not closed.
			if !isOpen{
				break
			}
			fmt.Println(a)
				time.Sleep(time.Millisecond * 1000)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Completed");
}