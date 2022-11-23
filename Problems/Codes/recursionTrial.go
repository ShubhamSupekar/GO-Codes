package main

import "fmt"

func main() {
	var a int = 100
	// fmt.Println("enter the number")
	// fmt.Scan(&a)
	// fmt.Println(recursion(a))
	b := recursion(a)
	fmt.Println(b)
}
func recursion(a int)(int) {
	if a > 10 {
		return recursion(a / 2)
	} 
	return a
}
