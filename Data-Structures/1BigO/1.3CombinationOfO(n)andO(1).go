package main

import (
	"fmt"
)

func main() {
	studentList := []string{"Soham", "Shubham", "Ganesh", "Sonal"} //O(1)
	a := []string{}                                                //O(1)
	increment := 0                                                 //O(1)
	for i := 0; i < len(studentList); i++ { 
		increment++                                                //O(n)
		a = append(a, studentList[i])                              //O(n)

	}

	fmt.Print("slice a is:", a)                                     //O(1)
	fmt.Print("\nslice studentlist is: ", studentList)              //O(1)
	fmt.Print("\nthe number of element in a slice is :", increment) //O(1)
} //O(6+2n) = calculation fo complexity
