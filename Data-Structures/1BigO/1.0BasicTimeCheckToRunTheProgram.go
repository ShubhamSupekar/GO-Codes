//calculate the sum of 100 natural numbers
package main

import (
	"fmt"
	"time"
)

func main() {
	now1 := time.Now()
	var sum int
	//var n int
	//fmt.Print("Enter a positive integer : ")
	//fmt.Scan(&n)
	for i := 1; i <= 100; i++ { // assigning 1 to i
		sum += i // sum = sum + i
	}
	fmt.Print("Sum : ", sum)
	fmt.Print("\nTime required to run a program: ", time.Since(now1))
}

// we acnnot depend on system clock because if you run program
// it will give different results again and again
//hence we requires BIG O
