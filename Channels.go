package main

import ("fmt";"time";"sync")

func main(){
	wg := sync.WaitGroup{}
	wg.Add(3)
	ch := make(chan int)
	go func(){
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

	data := make(chan int)
	go func(){
		for true{
			fmt.Println("Getting data from channel")
			a,isOpen := <- data
			if !isOpen {
				break
			}
			fmt.Println("Getting data: {}", a)
			time.Sleep(time.Millisecond*1000)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Completed");
}