package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i <= 10; i++ {
		wg.Add(2)
		m.RLock() // lock reader
		go SayHello()
		m.Lock() //writer lock
		go increment()
	}
	wg.Wait()
}
func SayHello() {
	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	m.Unlock()
	wg.Done() //this program is for printing 20 gorutines but in the output
	// all the numbers are messed up this is called a RACE condition
}
