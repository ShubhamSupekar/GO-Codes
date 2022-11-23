package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i <= 10; i++ {
		wg.Add(2)
		go SayHello()
		go increment()
	}
	wg.Wait()
}
func SayHello() {
	fmt.Printf("Hello %v\n", counter)
	wg.Done()
}
func increment() {
	counter++
	wg.Done() //this program is for printing 20 gorutines but in the output
	// all the numbers are messed up this is called a RACE condition
}
