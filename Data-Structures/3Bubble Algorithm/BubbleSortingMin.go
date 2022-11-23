package main

import (
	"fmt"
)

func main() {
	a := []int{100, 20, 10, 70, 1000}
	for j := 1; j < len(a); j++ {
		for i := 0; i < len(a)-j; i++ {
			if a[i] < a[i+1] {
				//swap
				temp := a[i]
				a[i] = a[i+1]
				a[i+1] = temp
				//continue
			}
			//fmt.Println(a[0])
		}
	}
	fmt.Println(a)
	fmt.Println("min value is: ", a[len(a)-1])
}
