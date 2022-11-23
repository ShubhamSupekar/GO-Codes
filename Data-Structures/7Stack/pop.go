package main

import "fmt"

func main() {
	stack := []int{10, 20, 30, 40, 50}
	var number int
	for {
		fmt.Println("enter 1 for pop (press q to stop):")
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println(stack)
			break
		}
		fmt.Println("you are poping:",stack[0])
		fmt.Println("before poping",stack)
		stack = append(stack[1:])
		fmt.Println("after poping stack looks like: ",stack)
	}
}
