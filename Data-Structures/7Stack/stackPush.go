package main

import "fmt"

func main() {
	fmt.Println("enter the legth of an array: ")
	var length int 
	fmt.Scan(&length)
	a := make([]int,length)
	var n int
	var i int = length
		if length==0{
			fmt.Println("this is an empty array",a)
		}else {
			for ;i>0;i--{
				fmt.Println("enter the number push for 1 or end for 2:")
				var decision int
				fmt.Scan(&decision)
				if decision ==1 {
					fmt.Println("enter the number: ")
					fmt.Scan(&n)
					a[i-1] =n
					continue
				}else {
					fmt.Println("This is the end")
						break
				}				
			}
			fmt.Println(a)
	}
}
