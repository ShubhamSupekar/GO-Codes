package main

import (
	"fmt"
)

// left root right
type node struct {
	data  int
	left  *node
	right *node
}

var result []int

func treeBuild(root *node, k int) {
	//building tree
	if root.data > k {
		if root.left == nil {
			root.left = &node{data: k}
			result = append(result, k)
			return
		}
		treeBuild(root.left, k)
	} else {
		if root.right == nil {
			root.right = &node{data: k}
			result = append(result, k)
			return
		}
		treeBuild(root.right, k)
	}
	return
}
func printing(root *node) []int {
	output := make([]int, 0)
	if root == nil {
		return nil
	}
	left := printing(root.left)
	right := printing(root.right)
	output = append(output, left...)
	output = append(output, root.data)
	output = append(output, right...)
	return output
}
func main() {
	var root *node
	a := []int{60, 80, 90, 10, 20, 50, 40, 200, 5}
	for _, b := range a {
		if root == nil {
			root = &node{data: b}
			result = append(result, b)
			continue
		}
		treeBuild(root, b)
	}
	fmt.Println(printing(root))
	fmt.Println("result is: ", result)
}
