package main

import "fmt"

func main() {
	var a = []int{10, 20, 20, 10, 10, 30, 50, 10, 20, 30}
	dataRepetation := map[int]int{}
	for i := 0; i < len(a); i++ {
		if dataRepetation[a[i]] == 0 {
			dataRepetation[a[i]] = 1
			continue
		} else {
			b, _ := dataRepetation[a[i]]
			b = b + 1
			dataRepetation[a[i]] = b
			continue
		}
	}
	var b int
	fmt.Println(dataRepetation)
	for _, d := range dataRepetation {
		b += d / 2
	}
	fmt.Println(b)
}
