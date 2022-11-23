	package main

import "fmt"

func main() {
	a := []int{500, 65, 99, 87, 74, 63, 76, 100}
	b:= []int{1 ,12 ,5 ,111 ,200 ,1000, 10}
	fmt.Println(devide(b),"||",devide(a))
}
func devide(array []int) []int {
	if len(array) < 2 {
		return array
	}
	first := devide(array[:len(array)/2])
	second := devide(array[len(array)/2:])
	return mergesort(first, second)
}
func mergesort(a, b []int) (result []int) {
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}
	return result
}
