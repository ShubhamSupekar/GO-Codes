package main

import (
	"fmt"
)

func main() {
Loop:
	for i := 1; i <= 9; i++ {
		fmt.Println(i)
		if i == 3 { // we stop this loop at 3
			break Loop
		}
	}
}
