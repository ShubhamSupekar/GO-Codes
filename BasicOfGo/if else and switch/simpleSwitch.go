package main

import (
	"fmt"
)

func main() {
	switch i := 0 + 2; i { // enter here 1 or 3 for next results
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	default:
		fmt.Println("not one or two")
	}
}
