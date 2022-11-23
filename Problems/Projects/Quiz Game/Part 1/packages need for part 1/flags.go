package main

import (
	"flag"
	"fmt"
)

func main() {
	word := flag.String("Word", "Foo", "a string")
	flag.Parse() //Once all flags are declared,
	//call flag.Parse() to execute the command-line parsing
	fmt.Println("Word:", *word)
}

//Use -h or --help flags to get automatically generated help text for the command-line program.
//PS S:\go workspace\Quiz Game\Part 1\packages need for part 1> go run .\flags.go -Word=opt
// output is: Word: opt
