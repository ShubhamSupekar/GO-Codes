package main

import "fmt"


func main(){
	a := []int{55,7,0,6,79,1,9,47,77,9,4,10}
		b := a[:len(a)-1]
		c := a[1:]
		fmt.Println(a,b,c)
}