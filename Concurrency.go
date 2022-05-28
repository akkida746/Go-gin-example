package main

import ( "fmt";"sync";"time")

func main() {
	fmt.Println("Concurrency")
	wg := sync.WaitGroup{}
	wg.Add(2)
	fmt.Println("Completed")
	go func(){
		fmt.Println("First worker")
		time.Sleep(time.Millisecond*1000)
		wg.Done()
	}()
	go func(){
		fmt.Println("Second worker")
		time.Sleep(time.Millisecond*1000)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Completed")
}
func worker() {
	fmt.Println("Inside worker()")
}
