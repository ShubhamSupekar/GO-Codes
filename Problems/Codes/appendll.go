package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {
	var head *node
	var last *node
	for{
		fmt.Println("enter the index fro replace")
		var f int
		fmt.Scan(&f)
		a := []int{10, 20, 500, 40, 50, 60}
		if f<len(a){
			fmt.Println("enter the data at that index")
			var d int
			fmt.Scan(&d)
			for i := 0; i < len(a); i++ {
				n := node{data: a[i]}
				if head == nil {
					head = &n
					last = head
					continue
				}
				last.next = &n
				last = last.next
				if i == f {
					last.data = d
					last.next = &n
					continue
				}
			}
			for head != nil {
				fmt.Println(head.data,head.next)
				head = head.next
			}
			fmt.Println(a)
			continue
		}else {
			fmt.Println("enter the valid number")
			break
		}
	}
}