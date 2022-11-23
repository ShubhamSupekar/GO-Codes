package main

import (
	"fmt"
)

type bill struct {
	flag    bool
	counter int64
	pi      float64
}

type eg struct {
	flag    bool
	counter int64
	pi      float64
}

func main() {
	var b1 bill

	e1 := eg{
		flag:    true,
		counter: 15,
		pi:      3.141592,
	}

	b1 = bill(e1)
	fmt.Println(b1)
}
