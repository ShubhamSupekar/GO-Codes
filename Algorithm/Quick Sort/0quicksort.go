package main

import "fmt"

func recursion(a []int, l, hi int) {
	if l >= hi || l < 0 {
		return
	}
	l, hi = sort(a, l, hi)
	recursion(a,l,6)
	return
}
func sort(a []int, low, hi int) (int, int) {
	pivot := low
	for {
		if low == pivot {
			low++
			continue
		}
		if a[low] > a[pivot] && a[hi] < a[pivot] {
			a[low], a[hi] = a[hi], a[low]
			hi--
			low++
		} else if a[low] < a[pivot] {
			low++
		} else if a[hi] > a[pivot] {
			hi--
		}
		if low > hi {
			a[pivot], a[hi] = a[hi], a[pivot]
			pivot = hi
			break
		}
	}
	//fmt.Println(pivot)
	return low, hi
}
func main() {
	a := []int{4, 6, 1, 5, 3, 7, 2}
	hi := len(a) - 1
	recursion(a, 0, hi)
	fmt.Println(a)
}
