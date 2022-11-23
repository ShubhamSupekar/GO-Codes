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
	b := Bird{
		Animal:   Animal{Name: "EMU", Origin: "Australia"},
		SpeedKph: 40,
		CanFly:   false,
	}
	fmt.Println(b) //we add name and origin from animal struct to bird
}
