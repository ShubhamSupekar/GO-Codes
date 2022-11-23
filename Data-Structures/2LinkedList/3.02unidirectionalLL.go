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
	}
	for head != nil {
		fmt.Println(head.data, head.next)
		head = head.next
	}
}
 
