package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}



func (n *node) tree(k int) {
		//building tree
	if n.data < k {
		//move right
		if n.right == nil {
			n.right = &node{data: k}
		} else {
			n.right.tree(k)
		}
	} else if n.data > k {
		//move left
		if n.left == nil {
			n.left = &node{data: k}
		} else {
			n.left.tree(k)
		}
	}
	fmt.Println("Left is ",n.left)
	fmt.Println("Right is ",n.right)
}

func main() {
	var head *node
	fmt.Println("enter head ")
	var i int
	fmt.Scan(&i)
	head = &node{data: i}
	for {
		fmt.Println("Enter a number in tree ")
		fmt.Println("press q to exit")
		_, err := fmt.Scan(&i)
		if err != nil {
			fmt.Println(head.left, head.right)
			break
		}
		head.tree(i)
		fmt.Println("left is ",head.left)
		fmt.Println("right is ",head.right)
		continue
	}
}
