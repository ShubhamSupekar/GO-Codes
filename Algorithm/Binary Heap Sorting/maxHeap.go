package main

import (
	"fmt"
	"math"
	"reflect"
)

type heap struct {
	h []int
}

func newHeap() *heap {
	return &heap{
		h: []int{int(math.Inf(1))},
	}
}

/*

// These methods are not needed right now.

func (h *heap) left(i int) int {
	l := 2 * i
	if l >= len(h.h) {
		return -1
	}

	return l
}

func (h *heap) right(i int) int {
	r := 2*i + 1
	if r >= len(h.h) {
		return -1
	}

	return r
}

*/

func (h *heap) parent(i int) int {
	p := int(i / 2)
	if p == 0 {
		return -1
	}

	return p
}

func (h *heap) insert(num int) {
	h.h = append(h.h, num)

	// heapify the latest element.
	h.heapify(len(h.h) - 1)
}

func (h *heap) heapify(idx int) {
	pIdx := h.parent(idx)
	if pIdx == -1 {
		return
	}

	// If parent is greater than child then just return
	if h.h[pIdx] > h.h[idx] {
		return
	}

	// Swap value otherwise
	h.h[pIdx], h.h[idx] = h.h[idx], h.h[pIdx]
	h.heapify(pIdx)
}

func main() {
	h := newHeap()

	input := []int{24, 4, 42, 86, 51, 1, 3, 62}
	for _, num := range input {
		h.insert(num)
	}

	outputWanted := []int{-9223372036854775808, 86, 62, 24, 51, 42, 1, 3, 4}
	outputCalc := h.h

	if !reflect.DeepEqual(outputCalc, outputWanted) {
		fmt.Println("not equal")
		fmt.Println("wanted  :", outputWanted)
		fmt.Println("received:", outputCalc)
		return
	}

	fmt.Println("equal")
	fmt.Println("output is: ",outputCalc)
}
