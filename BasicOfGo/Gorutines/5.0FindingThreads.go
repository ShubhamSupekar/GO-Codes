package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Threads is %v\n", runtime.GOMAXPROCS(-1))
}
