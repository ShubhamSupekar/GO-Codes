package main

import (
	"fmt"
)

func main() {                       //this is the way we can reduce the caode
	count := 10
	fmt.Print("Value of count is ", count, "\nAddress of count is ", &count)
}
