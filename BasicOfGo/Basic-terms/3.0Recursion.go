package main

import (
	"fmt"
)

func main(){
	array:= []int{5,57,70,0,6,79,1,9,47,77,94,10}
		c := recursion(array)
		fmt.Println("Final c is: ",c)
}

func recursion(a []int) []int{
	if len(a)<=1{
		return a
	}
	fmt.Println("a is ",a)
	return recursion(a[:len(a)-1])
}