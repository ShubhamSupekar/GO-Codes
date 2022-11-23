package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {
	var head *node
	var last *node
	a := []int{10, 20, 30, 40, 50, 60}
	
	for _, val := range a {
		n := node{
			data: val,
		}
		if head == nil {
			head = &n
			last = head
			continue
		}
		last.next = &n
		last = last.next
		last.next = head
	}
	for counter:=len(a)+1;head!=nil;counter--{	
		if counter>0{
		fmt.Println(head.data, head.next,counter)
		head = head.next
		}else {
			fmt.Println("Continue..........")
			break
		}
	}
}
 
