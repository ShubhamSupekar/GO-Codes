package main

import "fmt"

type node struct {
	data int
	next *node
}
var head *node = new(node)

func main() {
	a := []int{10, 20, 30, 40, 50, 60}
	
	for _, val := range a {
		n := node{
			data: val,
		}
		if head == nil {
			head = &n
			continue
		}
		last := head
		for last.next != nil {
			last = last.next
		}
		last.next = &n
	}
	for head != nil {
		fmt.Println(head.data, head.next)
		head = head.next
	}
}
 
