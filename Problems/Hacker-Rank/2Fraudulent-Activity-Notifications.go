package main

import "fmt"
//only 3 test pass
func main() {
	var d int32 = 4
	a := []int32{1, 2, 3, 4, 4,5,10,55,80,1,45, 9,66,12,0,00,1,50,05,48,33,45,1,54,84,84,651,1,45,614,5,64,44,8,9,44,5,61,1,86,4,964,5,61,4,16,48,64,65,16,56,486,486,486,418,641,485,641,68,641,8641,8,64,8,641,8,64,84,46,41}
	var notification int32
	var b []int32
	var i int32
	for i = 0; (d + i) < int32(len(a)); i++ {
		b = a[i : d+i]
		e := divide(b)
		median := FindingMedian(e)
		if float32(a[d+i]) >= 2*median {
			notification += 1
		} else {
			continue
		}
	}
	fmt.Println("number of notifications sent is: ", notification)
}
func divide(array []int32) []int32 { // dividing an array
	if len(array) < 2 {
		return array
	}
	first := divide(array[:len(array)/2])
	second := divide(array[len(array)/2:])
	return mergesort(first, second)
}

func mergesort(a, b []int32) (result []int32) { // sorting
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

func FindingMedian(b []int32) float32 { // finding median
	mid := len(b) / 2
	var median float32
	if len(b)%2 == 0 {
		median = float32((b[mid] + b[mid-1])) / 2
	} else {
		median = float32(b[mid])

	}
	return median
}
