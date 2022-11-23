package main

import (
	"fmt"
)

func main() {
	var b string = "Shubham"
	a := []string{"Soham", "Shubham", "Ganesh", "Sonal"} //no.of inputs=4
	for i := 0; i < len(a); i++ {
		if b == a[i] { // number of comperasion operation peform=4
			fmt.Print("Available") // no. of i/p = no. of operations
			// therefore complexity is O(n)
		}
	}
}

//this is an pattern of O(n) because
