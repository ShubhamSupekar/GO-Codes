	package main

import "fmt"
//First In First Out == FIFO
// Top is at the end of the array.

// 0, 1, 2, 3
//          top

// 0, 1, 2, 3, 4
//             top

func main() {
	var stack []int

	for length := 0; length >= 0; {
		fmt.Println("Stack looks like this:")
		fmt.Println(stack)
		fmt.Println("Enter the number to push on stack or")
		fmt.Println("press p to pop an element from stack:")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("you are poping:", stack[0])
			stack = append(stack[1:])
			length--
			if length == 0 {
				fmt.Println("can not pop any element after this because")
				fmt.Println("you pop out last element also now,")
				fmt.Println("stack looks like:", stack)
				break
			}
			continue
		}
		stack = append(stack,input)
		length = len(stack)
	}
}
