package main

import "fmt"

type node struct {
	data   int
	before *node
	next   *node
}

func main() {
	var head, storeN *node
	element := 100
	array := []int{1, 2, 3, 4, 5}
	for _, val := range array {
		n := &node{data: val}
		if head == nil {
			head = n
			storeN = n
			continue
		}
		n.before = storeN
		storeN.next = n
		storeN= n
		continue
	}
	store := head
	fmt.Println("LL Befor insert")
	for head != nil {
		fmt.Println(head.before, head.data, head.next)
		head = head.next
	}
	fmt.Println("_____________________")
	fmt.Printf("after appending element:%d in LL is:\n", element)
	store = insert(store, element)
	for store != nil {
		fmt.Println(store.before, store.data, store.next)
		store = store.next
	}
}

func insert(ll *node, a int) *node {
	var store1 *node
	store3 := ll //1 2 3 4 5     100
	t := &node{data: a}
	for {
		if ll.next == nil {             //5.next == nil 
			temp := store1.next         //temp = 5 
			store1.next = t //100       // store1.next = 100
			t.next = temp               //100.next =  5
			t.before = store1           // 100.before = 4
			temp.before = t             //5.before = 1000
			break
		}
		store1 = ll //4
		ll = ll.next //5
	}

	return store3
}
