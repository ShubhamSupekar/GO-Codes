package main

import (
	"fmt"
)

func main() {
	greeting := "Hello"
	name := "Soham"
	sayName(&greeting, &name)
	fmt.Println(name)
}
func sayName(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Shubham"
	fmt.Println(*name)
}
