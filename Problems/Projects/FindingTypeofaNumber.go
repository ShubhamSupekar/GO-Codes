package main

import (
	"fmt"
)

func main() {
	var i interface{} = "Shubham"
	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case float64:
		fmt.Println("i is float")
		break
		fmt.Println("HI") // breaking out early
	case string:
		fmt.Println("i is string")
	}
}
