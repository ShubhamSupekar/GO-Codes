package main

import (
	"fmt"
	"os"
    "io"
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
}
func print(w io.Writer, node *node, ns int, ch rune) {
    if node == nil {
        return
    }
 
    for i := 0; i < ns; i++ {
        fmt.Fprint(w, " ")
    }
    fmt.Fprintf(w, "%c:%v\n", ch, node.data)
    print(w, node.left, ns+2, 'L')
    print(w, node.right, ns+2, 'R')
}
func main() {
	var k int = 100
	var head *node
	for {
		if head == nil {
			head = &node{data: k}
			continue
		}
		// head.tree(100)
		head.tree(-20)
		head.tree(-50)
		head.tree(-15)
		head.tree(-60)
		head.tree(50)
		head.tree(60)
		head.tree(55)
		head.tree(85)
		head.tree(15)
		head.tree(5)
		head.tree(-10)
		fmt.Println(head.data)
		print(os.Stdout, head, 0, 'M')
		break
	}
}
