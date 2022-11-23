package main

import (
	"fmt"
)
func main(){
	a := []int{90,80,70,40,50,0,10}
	length := len(a)
	for i:=0;i<length;i++{
		store := a[i]
		position := i
		for j:=position;j>=0;j--{
			if store < a[j]{
				a[j+1]=a[j]
				a[j] = store
			}
			
		}

	}
	fmt.Println(a)
}