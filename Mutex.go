package main
import ("fmt";"sync")

//https://www.fareez.info/blog/thread-safe-lazy-loading-using-sync-once-in-go/

var mutex = sync.Mutex{}
var config = 0
func main(){
	fmt.Println("Starting")
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i:=1;i<=10;i++{
		go func(){
			loadConfig()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Completed")
}

func loadConfig() int {
	mutex.Lock()
	if config == 0{
		fmt.Println("Inside loadconfig")
		config = 1
	}
	mutex.Unlock()
	return config
}