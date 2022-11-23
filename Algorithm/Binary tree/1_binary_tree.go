package main

import (
	"fmt"
	"io"
	"os"
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
func prints(w io.Writer, node *node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	prints(w, node.left, ns+2, 'L')
	prints(w, node.right, ns+2, 'R')
}
func printing() {
	//fmt.Println(stack[0])
	if len(stack) == 0 {
		return
	}
	curr := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	//fmt.Println(stack)
	fmt.Printf("{")
	if curr.left != nil {
		stack = append(stack, curr.left)
		fmt.Printf(" L:%d,", curr.left.data)
	}
	fmt.Printf(" C:%d,", curr.data)
	if curr.right != nil {
		stack = append(stack, curr.right)
		fmt.Printf(" R:%d ", curr.right.data)
	}
	fmt.Println(" }")
	printing()
}
func main() {
	var k int = 6
	var head *node
	for {
		if head == nil {
			head = &node{data: k}
			continue
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
		printing()
		prints(os.Stdout, head, 0, 'M')
		break
	}
}
