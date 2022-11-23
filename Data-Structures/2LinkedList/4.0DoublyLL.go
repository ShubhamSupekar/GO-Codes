package main

import "fmt"

type node struct {
	data   int
	before *node
	next   *node
}

func main() {
	var head, storeN *node
	array := []int{1, 2, 3, 4, 5}
	for _, val := range array {
		n := &node{data: val}
		if head == nil {
			head = n
			storeN = n
			continue
		}
		n.before = storeN
		storeN.next = n
		storeN = n
		continue
	}
	for head != nil {
		fmt.Println(head.before, head.data, head.next)
		head = head.next
	}
}
