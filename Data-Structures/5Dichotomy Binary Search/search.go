//Find the index position of a given value from an already ordered array.
package main

import (
	"fmt"
)
func main(){
	a:= []int{10,20,30,40,50}
	mid := len(a)/2
	length := len(a)-1
	var n int
	fmt.Println("Enter the number to find in the array: ")
	fmt.Scan(&n)
	if n>a[length]{
		fmt.Println("enter the valid number")
	}else {
		
	searchValue := n
	for i:=1;i<=mid;i++{
		if searchValue>a[mid]{
			temp:=mid
			temp += i
			if a[temp]==searchValue{
				fmt.Println("value found at index: ",temp)
			}
		}
		if searchValue<a[mid]{
			temp := mid
			temp = temp - i
			if a[temp]==searchValue{
				fmt.Println("value found at index: ",temp)
			}
		}
		if searchValue==a[mid]{
			fmt.Println("value found at middle with an index: ",mid)
			break
		}

	}

	}
}