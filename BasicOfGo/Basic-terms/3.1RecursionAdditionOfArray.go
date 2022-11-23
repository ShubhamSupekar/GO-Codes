package main

import "fmt"

func recursion(a []int)int{
	if len(a)==0{
		return 0
	}	
	fmt.Println("len of a is: ",len(a))
	sum := recursion(a[:len(a)-1])+a[len(a)-1]
	return sum
} 

func main(){
	a := []int{1,2,3}
	fmt.Println(recursion(a))
}