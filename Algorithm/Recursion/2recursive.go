package main

//return function
import "fmt"

func main() {
	var n int
	for {
		fmt.Println("\nenter a number(press q to quit): ")
		_, err := fmt.Scan(&n)
		if err == nil {
			result := recursive(n)
			fmt.Printf("factorial of %d is %d", n, result)
		} else {
			fmt.Println("END")
			break
		}
	}
}
func recursive(a int) int {
	if a == 1 {
		return 1
	}
	return recursive(a-1) * a
}
