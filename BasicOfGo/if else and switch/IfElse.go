package main

import (
	"fmt"
)

func main() {
	a := 40
	b := 30
	if a < 1 || b > 100 {
		fmt.Println("number must be in between 1 and 100")
	} else {
		fmt.Println("number is between 1 and 100")
	}
	fmt.Println(a >= b, a <= b, a == b, a != b)
}
