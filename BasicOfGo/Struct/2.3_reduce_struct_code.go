package main

import (
	"fmt"
)

func main() { //this is the way we can reduce the caode
	e := struct {
		flag    bool
		counter int64
		pi      float64
	}{
		flag:    true,
		counter: 16,
		pi:      3.141592,
	}
	fmt.Println(e)
}
