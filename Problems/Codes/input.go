package main

import "fmt"
type node struct{
	data int
	next *node
}

func main() {
	var s int
	
	for{ 
		fmt.Println("enter the number: ")
		fmt.Scan(&s)
		
		if s == 0{
		fmt.Println("enter number greater than zero")
		break
		}else {
			ll(s)
			continue
		}
	}	
}
func ll(s int){
	var head *node
	var last *node
	n:=node{data:s,next: nil}
		for head == nil{
		head = &n
		head = head.next
		last = head
		break
		}
		last.next = &n
		last = last.next
		for head != nil{
			fmt.Println(head.data,head.next)
			head = head.next
		}
}