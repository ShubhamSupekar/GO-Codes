package main

import (
	"fmt"
)

func main() {
	a := 110
	defer fmt.Println(a)
	a = 21
}
