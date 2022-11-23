package main

import (
	"fmt"
)

func main() {
	total := 0 //O(1)

	a := []int{1, 2, 3, 4}
	for i := 0; i < len(a); i++ {
		//fmt.Printf("\ntotal is %d and element is %d", total, a[i])
		for j := 0; j < len(a); j++ {
			fmt.Printf("\n %d,%d", a[i], a[j]) //O(n^2)
			total++                            // O(n^2)
		}

	}
	fmt.Println("\ntotal process is: ", total) //O(1)
	// total is O(2+2(n^2)) = O(n^2)
}
