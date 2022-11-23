package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	fmt.Println("root value is: ", root.Val)
	if root.Left == nil && root.Right == nil {
		return 1
	}
	lHeight := maxDepth(root.Left)
	fmt.Println("left part: ", lHeight)
	rHeight := maxDepth(root.Right)
	fmt.Println("Right part: ", rHeight)
	if lHeight >= rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1
	}
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	root.Right.Right.Right = &TreeNode{Val: 16}
	height := maxDepth(&root)
	fmt.Println(height)
}
