package main

import (
	"fmt"
)

// but there is an error in this method
func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a) // see this output for error

}
