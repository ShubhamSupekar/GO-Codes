// by suraj dada
package main

import "fmt"

// Top is always at the start of the array.

// 1
// top

// 2, 1
// top

// 3, 2, 1
// top

func main() {
	var stack []int

	for {
		fmt.Println("Enter the number to push on stack (press q to stop):")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}

		stack = append([]int{input}, stack...)
	}

	fmt.Println("Stack looks like this:")
	fmt.Println(stack)
}