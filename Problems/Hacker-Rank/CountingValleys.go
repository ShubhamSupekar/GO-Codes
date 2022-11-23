package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "DDUUDDUDUUUD"
	var counter, Vallycount int32
	var seaLevel = true
	array := strings.Split(path, "")
	for _, c := range array {
		if c == "U" {
			counter++
		} else {
			counter--
		}
		if seaLevel && counter == -1 {
			Vallycount += 1
		}
		seaLevel = counter == 0
	}
	fmt.Println(Vallycount)
}

// "DDUUDDUDUUUD"

//                  _     _      /\_
//                    \  / \    /
//                     \/   \/\/
