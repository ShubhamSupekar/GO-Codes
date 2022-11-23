package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	var start *node
    var last *node
	// generate linked list
	for _, num := range nums {
		n := node{
			data: num,
		}
		if start == nil {
			start = &n
			continue
		}
		last = start
		for last.next != nil {
			last = last.next
		}
		last.next = &n
	}
    last.next=start

	// print LL
	now := start
	for now != nil {
		fmt.Println(now.data,now.next)
		now = now.next
	}
}