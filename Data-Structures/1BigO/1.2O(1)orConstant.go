package main

import (
	"fmt"
)

func main() {
	studentList := []string{"Soham", "Shubham", "Ganesh", "Sonal"} //no.of inputs=4
	fmt.Println(studentList[0])
}

//it dosent matter how many i/p there
//we only want to print first name in list
//therefore the complexity is constant or O(1)
//if we add fmt.Println(studentList[1]) then also not caring
//about number of inputs we only want to display 2nd name
