package main

import (
	"fmt"
)

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("the sum is: ", s)
}
func sum(values ...int) (result int) {
	fmt.Println(values)
	result = 0
	for _, v := range values {
		result += v
	}
	return

}
