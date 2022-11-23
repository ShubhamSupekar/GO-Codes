package main

import (
	"fmt"
)

func main() {
	a := []int{} //an empty array
	fmt.Println(a)
	fmt.Printf("Length = %v\n and capacity = %v\n", len(a), cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length = %v\n and capacity = %v\n", len(a), cap(a))
	a = append(a, 2, 3, 4, 5, 6)
	fmt.Println(a)
	fmt.Printf("Length = %v\n and capacity = %v\n", len(a), cap(a))

}
