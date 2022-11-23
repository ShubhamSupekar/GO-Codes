package main

import "fmt"

type node struct {
	data  int
	right *node
}

var a = make([]int, 10)

func Creat(b int) {
	idx := b % 10
	if a[idx] != 0 {
		collision(idx, b)
		return
	}
	a[idx] = b
}
func find(c int) {
	idx := c % 10
	if a[idx] == c {
		fmt.Println(c, "value exist")
		return
	}
	fmt.Println("value does not found")
}
func collision(idx, b int) {
	var head *node
	head = &node{data: a[idx]}
	head.right = &node{data: b}
	fmt.Printf("{")
	for head != nil {
		//fmt.Println("there is an collision")
		fmt.Printf(" %d ", head.data)
		head = head.right
	}
	fmt.Printf("}")
	return
}
func main() {
	Creat(15)
	Creat(12)
	Creat(10)
	Creat(11)
	Creat(21)
	Creat(31)
	Creat(25)
	Creat(13)
	Creat(19)
	Creat(16)
	fmt.Println(a)
	find(19)
	find(22)
	find(25)
}
