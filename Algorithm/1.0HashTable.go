package main

import "fmt"

type node struct {
	data int
	next *node
}

type hashTable struct {
	table []*node
}

var hashMod = 10

func CreateHashTable() *hashTable {
	return &hashTable{
		table: make([]*node, hashMod),
	}
}

func hash(k int) int {
	return k % hashMod
}

func (h *hashTable) put(num int) {
	idx := hash(num)

	n := node{data: num}
	if h.table[idx] == nil {
		h.table[idx] = &n
		return
	}

	curr := h.table[idx]
	var prev *node
	for curr != nil {
		// The number already exists in the hashtable we don't need to add it again.
		if curr.data == num {
			return
		}

		prev = curr
		curr = curr.next
	}

	prev.next = &n
}

func (h *hashTable) remove(num int) {
	idx := hash(num)

	if h.table[idx] == nil {
		return
	}

	curr := h.table[idx]
	var prev *node
	for curr != nil {
		// This is the first element in the list and we wanna remove it.
		if curr.data == num && prev == nil {
			h.table[idx] = curr.next
			return
		}

		if curr.data == num {
			prev.next = curr.next
			return
		}

		prev = curr
		curr = curr.next
	}
}

func (h *hashTable) get(num int) bool {
	idx := hash(num)

	curr := h.table[idx]
	for curr != nil {
		if curr.data == num {
			return true
		}

		curr = curr.next
	}

	return false
}

func (h *hashTable) String() string {
	var ret string

	for idx, n := range h.table {
		ret += fmt.Sprintf("%d:\t", idx)

		curr := n
		for curr != nil {
			ret += fmt.Sprintf("%d,", curr.data)
			curr = curr.next
		}

		ret += "\n"

	}

	return ret
}

func main() {
	input := []int{24, 97, 15, 11, 81, 4, 19, 76, 10, 23, 27, 3, 10, 83, 52, 30, 37, 69, 67, 25, 34, 6}

	h := CreateHashTable()

	for _, n := range input {
		h.put(n)
	}

	fmt.Println(h)
	h.remove(88)
	fmt.Println(h)
}
