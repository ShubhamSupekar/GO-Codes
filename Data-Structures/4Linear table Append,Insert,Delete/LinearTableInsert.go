package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 6, 5, 8}
	length := len(a)
	temp := make([]int, length+1)
	n := 2 //enter the index you want to insert a number 
	m := 75 // enter the number
	for i := 0; i < length+1; i++ {
		if i < n {
			temp[i] = a[i]

		}
		if i == n {
			temp[i] = m

		}
		if i > n {
			temp[i] = a[i-1]
		}
		fmt.Println("temprary array ", temp[i])
	}
	a = temp
	fmt.Println("after the insertion array is ",a)

}
