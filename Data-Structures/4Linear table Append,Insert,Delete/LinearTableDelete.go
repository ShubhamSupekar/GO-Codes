package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 5, 8, 10}
	fmt.Println(a)
	length := len(a)
	temp := make([]int, length-1)
	
	var n int  //index of delete the element in the array
	fmt.Printf("Enter the index of the number to delete: ")
	fmt.Scan(&n)
	if n > length-1{
		fmt.Println("enter valid index")
	}else{
		for i := 0; i < length; i++ {
			if i < n {
				temp[i] = a[i]
			}
			if i > n {
				temp[i-1] = a[i]
			}
		
		}
		fmt.Println("before: ", a, len(a))
		fmt.Println(temp)
		a = temp
		fmt.Println("after: ", a, len(a))
	}
}
