package main

import (
	"fmt"
)

//https://www.youtube.com/watch?v=zLnJcAt1aKs

const ArraySize = 7

// Hash Table structure
type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket structure (LL)
type bucket struct {
	head *bucketNode
}

// Busket Node structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert  in HT
// func (h *HashTable) Insert(key string) {
// 	index := hash(key)
// 	h.array[index].insert(key)
// }

// serach  in HT
// func (h *HashTable) Serach(key string) bool {

// }

// Delete  in HT
// func (h *HashTable) Search(key string) {
// index := hash(key)

// }

// Hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Insert  in ll
func (b *bucket) Insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

// serach  in ll
func (b *bucket) serach(){

}
// Delete  in ll
// Init will create a bucket(LL) in each slot of hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHT := Init()
	fmt.Println(testHT)
	fmt.Println(hash("RANDY"))

	testBucket := &bucket{}
	testBucket.Insert("RANDY")
}
