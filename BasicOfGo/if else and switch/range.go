package main

import (
	"fmt"
)

func main() {
	s := []int{10, 20, 30}
	j := "Hello World"
	for k, v := range s {
		fmt.Println(k, v)
	}
	for i, k := range j {
		fmt.Println(i, string(k))
	}
}
