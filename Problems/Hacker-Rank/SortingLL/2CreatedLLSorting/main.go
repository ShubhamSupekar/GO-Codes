package main

import "fmt"

type node struct {
	data   int
	before *node
	next   *node
}

func main() {
	var head, storeN *node
	//element := 100
	array := []int{5, 3, 1, 4, 2, 6}
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
	store := head
	fmt.Println("____________BEFORE sorting___________")
	for head != nil {
		fmt.Println(head.before, head.data, head.next)
		head = head.next
	}
	fmt.Println("_____________________")
	fmt.Println("__________AFTER sorting the ll____________")
	head = sort(store)
	for head != nil {
		fmt.Println(head.before, head.data, head.next)
		head = head.next
	}
}

func sort(head1 *node) *node {
	
}
