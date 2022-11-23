package main

import (
	"fmt"
)

func Findingmedian(d int) int{
	//finding median
	var median int
	fmt.Println(d%2)
	if d%2 ==0{
		median = d/2
	}else {
		median = (d+1)/2
	}
	return median
}

func main(){
	var d int
	//a:=[]int{10,20,30,40}
	d = 3
	c:= Findingmedian(d)
	fmt.Println(c) 
}