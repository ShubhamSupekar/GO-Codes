package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}
type Bird struct {
	Animal   //Padding
	SpeedKph float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "EMU"
	b.Origin = "Australia"
	b.SpeedKph = 40
	b.CanFly = false

	fmt.Println(b) //we add name and origin from animal struct to bird
}
