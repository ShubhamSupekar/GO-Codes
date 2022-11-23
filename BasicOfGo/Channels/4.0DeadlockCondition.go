package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} //multiple senders and one receiver

func main() {
	ch := make(chan int)

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(2)
			i := 42
			ch <- i
			wg.Done()
		}()
	}
	wg.Wait()
}
