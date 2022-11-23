package main

import "fmt"

func BuyToys(p []int, budget int) int {
	var sum int
	// var c int
	e := divide(p)
	fmt.Println("After sorting aray looks like: ", e)
	var ie int
	for ie = 0; ie < len(e); ie++ {
		sum += e[ie]
		fmt.Println("sum is: ", sum)
		if sum == budget {
			break
		}	
		if sum > budget {
			break
		}
	}
	return ie
}

func divide(array []int) []int { // dividing an array
	if len(array) < 2 {
		return array
	}
	first := divide(array[:len(array)/2])
	second := divide(array[len(array)/2:])
	return mergesort(first, second)
}

func mergesort(a, b []int) (result []int) { // sorting
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

func main() {
	prices := []int{1, 12, 5, 111, 200, 1000, 10}
	budget := 50
	c := BuyToys(prices, budget)
	fmt.Printf("Mark can buy %d toys with %d budget", c, budget)
}
