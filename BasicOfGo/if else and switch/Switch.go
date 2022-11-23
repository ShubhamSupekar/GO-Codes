package main

import (
	"fmt"
)

func main() {
	i := 11
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
	case i > 10:
		fmt.Println("i is greater than 10")
	}
}
