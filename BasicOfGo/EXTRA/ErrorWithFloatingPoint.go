package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0.123
	if a == math.Pow(math.Sqrt(a), 2) { //squaring the number

		fmt.Println("they are the same")
	} else {
		fmt.Println("they are different")
	}
}

// floating point takes approximation of the numbers
