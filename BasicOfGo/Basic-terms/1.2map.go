package main

import (
	"fmt"
)

func main() {
	statePopulaion := map[string]int{
		"MH": 10203040,
		"GJ": 20304050,
		"UP": 30405060,
		"AP": 40506070,
	}
	fmt.Println(statePopulaion)
	statePopulaion["TN"] = 60708090
	fmt.Println(statePopulaion)
	delete(statePopulaion, "UP")
	fmt.Println(statePopulaion)
	fmt.Println(statePopulaion["UP"])
	_, ok := statePopulaion["CH"]
	fmt.Println(ok)
	a, check := statePopulaion["MH"] // a is for the value of MH and check to check mh is in the map or not
	fmt.Println(a, check)
	//if this output is 0 and false then element is not in the map
	fmt.Println(len(statePopulaion))

}
