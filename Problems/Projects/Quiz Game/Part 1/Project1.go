package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file with questions and answers")
	flag.Parse()
	// next we want to read this file
	file, err := os.Open(*csvFilename) //file open
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file %s\n", *csvFilename))
	}
	//once file is open we want to creat a reader
	r := csv.NewReader(file)
	lines, err := r.ReadAll() //read all the lines of csv
	if err != nil {
		exit("Failed to parse the provided csv file")
	}
	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string //store answer
		fmt.Scanf("%s\n", &answer)
		if answer == p.a { //checking answer is correct or not
			correct++
		}
	}
	fmt.Printf("Your score is %d out of %d\n", correct, len(problems))
}
func parseLines(lines [][]string) []problem { //2D string slice and return slice problems
	ret := make([]problem, len(lines)) //declear a variable to return
	for i, lines := range lines {      //creat slice with containing struct
		ret[i] = problem{
			q: lines[0],
			a: strings.TrimSpace(lines[1]), //trim space does not consier spaces before answers
		}
	}
	return ret
}

type problem struct {
	q string //for question
	a string //for answers
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

//go run .\Project1.go -csv="S:\go workspace\Quiz Game\Part 1\problems.csv"       :input
