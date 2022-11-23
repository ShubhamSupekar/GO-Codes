package main

import "fmt"

type node struct {
	data   int
	next  *node
	prev *node
}

func main() {
	var head *node
	var before *node
	var after *node
	a := []int{10, 20, 30, 40, 50, 60}
	for _, val := range a {
		n := node{data: val}
		for head == nil {
			head = &n
			after = head
			before = head
			continue
		}
		temp := &n
		temp.prev = before
		before = &n
		after.next = &n
		after = after.next
		if after.next == nil{ 
			after.next = head
			head.prev = after
		}
	}
	for counterrr := len(a); head != nil; counterrr-- {
		if counterrr > 0 {
			fmt.Println(head.prev,head.data,head.next,counterrr)
			head = head.next
		}else {
			fmt.Println("Continue......")
			break
		}
	}
}