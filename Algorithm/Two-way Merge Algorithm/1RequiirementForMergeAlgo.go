package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50, 60}
	divide(a)
}
func divide(array []int) []int {
	fmt.Println(array)
	if len(array) < 2 {
		return array
	}
	array = divide(append(array[:len(array)/2]))
	return array
}
