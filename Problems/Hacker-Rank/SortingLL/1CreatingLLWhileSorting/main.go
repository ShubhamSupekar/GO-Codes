package main

import "fmt"

type node struct {
	data   int
	before *node
	next   *node
}

func main() {
	var head , storeN*node
	arr := []int{22, 5255, 10, 31, 14, 6,0}
	//store the in ll and sort  it
	for _, val := range arr {
		n := &node{data: val}
		if head == nil{
			head = n
			storeN = n
			continue
		}
		storeN = head
		for n.data >storeN.data {
		//	temp1 = storeN
			if storeN.next == nil{
				storeN.next = n
				n.before= storeN
				break
			}
			storeN = storeN.next
			continue
		}
		if n.data < storeN.data{
			if storeN.before==nil{
				storeN.before = n
				n.next = storeN
				head = n
				continue
			}
			temp2 := storeN.before
			temp2.next = n
			n.before = temp2
			n.next = storeN
			storeN.before = n
			continue
		}
	}
	for head!=nil{
		fmt.Println(head.before,head.data,head.next)
		head = head.next
	}
}
