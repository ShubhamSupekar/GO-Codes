package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

var head, last *node

func insert(prev *node, p, v int) *node {
	var counter int = 1
	first := prev
	t := &node{data: v} // creating new node
	for prev.next != nil {
		if counter == p {
			temp := prev.next
			prev.next = t
			t.next = temp
		}
		prev = prev.next
		counter++
		continue
	}
	return first
}
func main() {
	position := 4
	val := 800
	array := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for _, c := range array {
		n := &node{data: c}
		if head == nil {
			head = n
			continue
		}
		last = head
		for last.next != nil {
			last = last.next
		}
		last.next = n
	}
	store := head
	fmt.Println("Before appending element in LL")
	for head != nil {
		fmt.Println(head.data, head.next)
		head = head.next
	}
	latst := insert(store, position, val)
	fmt.Println("__________________")
	fmt.Printf("After appending element: %d in position:%d LL \n", val, position)
	for latst != nil {
		fmt.Println(latst.data, latst.next)
		latst = latst.next
	}
}
