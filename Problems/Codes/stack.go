package main

import "fmt"

// Top is at the end of the array.

// 0, 1, 2, 3
//          top

// 0, 1, 2, 3, 4
//             top

func main() {
	var stack []int

	for {
		fmt.Println("Enter the number to push on stack (press q to stop):")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}

		stack = append(stack, input)
	}

	fmt.Println("Stack looks like this:")
	fmt.Println(stack)
}
