package main

import (
	"fmt"
	"io"
	"os"
)

type node struct {
	data  int
	left  *node
	right *node
}

//center left right     ==  Pre-order
func buildTree(curr *node, num int) {
	if curr.data < num {
		if curr.right == nil {
			curr.right = &node{data: num}
			return
		}

		buildTree(curr.right, num)
	} else {
		if curr.left == nil {
			curr.left = &node{data: num}
			return
		}

		buildTree(curr.left, num)
	}
}

var queue []*node

func print() {
	if len(queue) == 0 {
		return
	}

	// accessing data
	curr := queue[0]
	queue = queue[1:]

	fmt.Printf("{C:%d", curr.data)

	if curr.left != nil {
		// enqueing
		queue = append(queue, curr.left)
		fmt.Printf(", L:%d", curr.left.data)
	}

	if curr.right != nil {
		// enqueing
		queue = append(queue, curr.right)
		fmt.Printf(", R:%d", curr.right.data)
	}
	fmt.Println("}")
	print()
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

func main() {
	nums := []int{15, 28, 13, 3}
	root := &node{data: nums[0]}

	for i := 1; i < len(nums); i++ {
		buildTree(root, nums[i])
	}

	queue = []*node{root}
	fmt.Println()
	print()
	fmt.Println()
	prints(os.Stdout, root, 0, 'M')

}
