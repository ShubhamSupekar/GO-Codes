package main

import (
	"fmt"
)

func main() {
	a := []int{20, 40, 60, 80, 90}
	var n int
	fmt.Printf("Enter the number to search: ")
	fmt.Scan(&n)
	var search = false
	for i := 0; i < len(a); i++ {
		if n == a[i] {
			search = true
			fmt.Println("found the value in the array: ", n)
			fmt.Println("at the index: ", i)
		}
		break
	}
	if search != true{
		fmt.Println("value is not in the array")
	}
	
}
