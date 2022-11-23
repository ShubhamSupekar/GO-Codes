package main

import "fmt"

type Medical struct {
	SrNo      int
	Admin     string
	Coworkers []string
}

func main() {
	aMedical := Medical{
		SrNo:  1,
		Admin: "Shubham",
		Coworkers: []string{
			"Soham", "Siddhi", "Monu",
		},
	}
	fmt.Println(aMedical)
	fmt.Println(aMedical.Coworkers)

}
