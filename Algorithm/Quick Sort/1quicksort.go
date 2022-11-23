package main

import "fmt"

func recursion(a []int, l, hi int) {
	if l == 6 || hi == 0 {
		return
	}
	sort(a, l, hi)
	if l <=4 {
		l++
		fmt.Println("worked1",a,l)
		recursion(a,l,hi)
		return
	} 
	if l>4{
		l++
		fmt.Println("worked2",a)
		recursion(a,l,hi)
		return
	}

}

func sort(a []int, low, hi int) {
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
}
func main() {
	a := []int{4, 6, 1, 5, 3, 7, 2}
	// hi := len(a) - 1
	recursion(a, 0, 6)
	fmt.Println(a)
}
