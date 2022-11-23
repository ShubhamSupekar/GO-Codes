package main

import (
	"fmt"
)

func main() {
	data := []string{"ab", "ab", "bc"}
	queries := []string{"ab", "bc", "cd"}
	c := make(map[string]int)
	d := []int{}
	for _, b := range queries {
		c[b] = 0
	}
	for _, d := range data {
		a, check := c[d]
		if check == false {
			continue
		} else {
			a += 1
			c[d] = a
		}
	}
	fmt.Println(c)
	for _, b := range queries {
		f := c[b]
		d = append(d, f)
	}
	fmt.Println(d)
}
