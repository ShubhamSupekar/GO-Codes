package main

import (
	"fmt"
)

//building inorder : left root right
var stack []*node

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
		}
		n.right.tree(k)

	} else if n.data > k {
		//move left
		if n.left == nil {
			n.left = &node{data: k}
		}
		n.left.tree(k)

	}
}
func printing(rs *node) {
	if len(stack) == 0 {
		return
	}
	// fmt.Println(len(stack))
	curr := stack[0]
	stack = append(stack[1:])
	//fmt.Println(stack)
	if curr.left != nil {
		stack = append(stack,curr.left)
		fmt.Println()
		fmt.Printf("{L:%d,", curr.left.data)
		fmt.Printf("C:%d}", curr.data)
		//1st left
		printing(curr.left)
	}
	if curr.right != nil {
		stack = append(stack,curr.right)
		fmt.Println()
		fmt.Printf("{C:%d,", curr.data)
		fmt.Printf("R:%d}", curr.right.data)
		//2nd right print
		printing(curr.right)
	}
}
func main() {
	var k int = 6
	var head *node

	if head == nil {
		head = &node{data: k}
	}
	head.tree(8)
	head.tree(4)
	head.tree(3)
	head.tree(7)
	head.tree(12)
	head.tree(5)
	head.tree(9)
	head.tree(1)
	stack = []*node{head}
	printing(head)

}
