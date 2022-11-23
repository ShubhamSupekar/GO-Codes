package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.Greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) Greet() {
	fmt.Println(g.greeting, g.name)
}
