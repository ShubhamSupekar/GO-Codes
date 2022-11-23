package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "aba"
	b := 10
	var counter1 int
	var counter2 int
	var result int
	array := strings.Split(a, "")
	length := len(array)
	remainder := b % length
	for _, d := range array {
		if d == "a" {
			counter1++
		} else {
			continue
		}
	}
	if remainder == 0 {
		result = b * counter1
	} else {
		for i := 0; i <= remainder; i++ {
			if array[i] == "a" {
				counter2++
			}
		}
		result = (b*counter1) + counter2
	}
	fmt.Println(result)
}
