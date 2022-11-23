package main

import "fmt"

func main() {
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* output each array element's value */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Print(a[i][j]," ")
		}
		fmt.Println(" ")
	}

}
