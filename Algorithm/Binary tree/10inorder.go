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
func printing(root *node) []int {
	if root == nil{
		return nil
	}
	left := printing(root.left) 
	right := printing(root.right)
	output := make([]int,0)
	output = append(output, left...)
	output = append(output, root.data)
	output = append(output, right...)
	return output
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
	fmt.Println(printing(head))

}
