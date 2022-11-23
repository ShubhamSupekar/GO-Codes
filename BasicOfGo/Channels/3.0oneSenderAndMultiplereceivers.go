package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() { //receiver
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		go func() { // sender
			i := 42
			ch <- i
			wg.Done()
		}()
		wg.Wait()
	}
}
