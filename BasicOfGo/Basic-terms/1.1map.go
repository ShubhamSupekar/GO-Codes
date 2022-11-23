package main

import "fmt"

func main(){
	b := make(map [int]int)
	a := []int{10,20,30,10,20,50,80,50}
	for e,f := range a {
		if b[f] == 0{
			b[f]= 0
		}else {
			continue
		}
		fmt.Println(e)
	}
	fmt.Println(b)
}