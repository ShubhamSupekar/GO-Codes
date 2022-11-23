package main

import (
	"fmt"
	"strconv"
)

func main() {
	t := "07:05:45PM"
	s := string(t[8])
	if s == "A" {
		for j := 0; j < 8; j++ {
			fmt.Print(string(t[j]))
		}
	} else {
		m, _ := strconv.Atoi(string(t[0]))
		n, _ := strconv.Atoi(string(t[1]))
		// time := m*10 + n + 12
		// fmt.Print(time)
		// for i := 2; i < 8; i++ {
		// 	fmt.Print(string(t[i]))
		// }
		t[0] = (m+1)
		t[1] = (n+2)
		
	}
}
