package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads is %v\n", runtime.GOMAXPROCS(-1))
}
