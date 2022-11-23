package main

import (
	"fmt"
)

func main() {
	a := []int{1000, 40, 100, 20, 60}
	for n := 1; n < len(a)-1; n++ { //for multiple sorting to get precision result
		for i := 0; i < len(a)-n; i++ { //for finding max number only
			if a[i] > a[i+1] {
				temp := a[i]
				a[i] = a[i+1]
				a[i+1] = temp
			}
		}
		fmt.Println("max number is:", a[len(a)-1])
		fmt.Println(a)
	}
	fmt.Println("final value is:", a)
}
