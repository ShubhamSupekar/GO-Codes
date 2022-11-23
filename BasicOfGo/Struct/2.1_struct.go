package main

import (
	"fmt"
)

type eg struct { // we define a common structure fo all the users here
	flag    bool
	counter int64
	pi      float64
}

func main() { // we added different users for upper structure
	var e1 eg // var sets all values to 0
	e2 := eg{
		flag:    true,
		counter: 10,
		pi:      3.141259,
	}
	fmt.Print("this is an e1 values", e1)
	fmt.Print(" this is an e2 values", e2)
	fmt.Print(" \nthis is an e1 values", e1,"\nthis is an e2 values", e2)


}
