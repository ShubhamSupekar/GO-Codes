package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {
	// TODO: figure out a way to get data dynamically.
	nums := []int{1, 2, 3, 4, 5}

	var start *node

	// generate linked list
	for _, num := range nums {
		n := node{
			data: num,
		}

		// if start is nil, which means this is the first object of the list.
		if start == nil {
			start = &n
			continue
		}

		// find the current last node of the list.
		last := start
		for last.next != nil {
			last = last.next
		}
		last.next = &n
	}
	

	// print LL
	now := start
	for now != nil {
		fmt.Println(now.data,now.next)
		now = now.next
	}
}