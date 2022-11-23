package main

import (
	"fmt"
)

func main() {
	//var sum int
	//a := []int{1, 2, 1, 4}
	//for i := 0; i < len(a); i++ {
	//	sum += a[i]
	//}
	//fmt.Print(sum)
	a(1, 2, 3, 4, 5, 1)
}

func a(value ...int) {
	fmt.Print(value)
	sum := 0
	for i := 0; i < len(value); i++ {
		sum += value[i]

	}
	fmt.Print("\n", sum)

}
