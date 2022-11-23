package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} //multiple senders and one receiver

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) //closing channel to avoid dedlock condition
		wg.Done()
	}(ch)
	wg.Wait()
}
