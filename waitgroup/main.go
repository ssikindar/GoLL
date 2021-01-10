package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, waitgroup *sync.WaitGroup) {
	fmt.Println("kicking Goroutine \n", i)
	time.Sleep(2  * time.Second)
    fmt.Printf("Finishing %d \n", i)
	waitgroup.Done()
}

func main(){
	var waitgroup sync.WaitGroup
	for i:=0; i<5;i++ {
		waitgroup.Add(1)
		go process(i, &waitgroup)
	}
	waitgroup.Wait()
	fmt.Println("All go routines finished executing")

}