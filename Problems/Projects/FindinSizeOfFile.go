package main

import (
	"fmt"
)

const (
	kb = 1 << (10 * iota)
	mb
	gb
	tb
)

func main() {
	var FileSize float32 = 4000000
	fmt.Printf("file size in GB is %v", FileSize/gb)
}
