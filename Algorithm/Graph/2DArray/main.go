package main

import "fmt"

func main() {
	var a = [5][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}
	// here 5 rows and 2 column
	// (0,0) (0,1)
	// (1,0) (1,1) ......
	// (4,0) (4,1)

	fmt.Println("\nMatrix is: \n")

	for i := 0; i < 5; i++ {
		fmt.Print("( ")
		for j := 0; j < 2; j++ {

			fmt.Printf("%d ", a[i][j])

		}
		fmt.Println(")")
		fmt.Println("\n")

	}
}
