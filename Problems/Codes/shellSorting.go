package main

import "fmt"

func main() {
	a := []int{9, 6, 5, 8, 0, 7, 4, 3, 1, 2}
	length := len(a)
	counter := 4
	for mini := length / 2; mini > 0;counter-- {
		mid := (len(a)/2)-1
		for i := 0; i < mini; i++ {
			gap := mini
			if a[i] > a[mid+gap-i] {
				temp := a[mid+gap-i]
				a[mid+gap-i] = a[i]
				a[i] = temp
			}
		}
		length = mini
		mini = length/2
	}
	fmt.Println(a)
}
