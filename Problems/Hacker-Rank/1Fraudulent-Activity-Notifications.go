package main
//only 3 test pass
import "fmt"

var notification int

func elementing(a []int) float32 {
	b := make(map[int]int) // creating map
	for _, ef := range a { //elementing duplicate numbers for finding median
		if b[ef] == 0 {
			b[ef] = 0
		} else {
			continue
		}
	}
	fmt.Println(b)
	var c []int
	for fd := range b {
		c = append(c, fd) //appending into an array
	}
	fmt.Println("before sorting: ", c)
	c = divide(c)
	fmt.Println("after sorting: ", c)
	median := FindingMedian(c) //passing sorted array into function to find median
	return median
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
func FindingMedian(b []int) float32 { // finding median
	mid := len(b) / 2
	var median float32
	if len(b)%2 == 0 {
		median = float32((b[mid] + b[mid-1])) / 2
	} else {
		median = float32(b[mid])

	}
	fmt.Println(median)
	return median

}

func main() {
	d := 5
	a := []int{2, 3, 4, 2, 3, 6, 8, 4, 5}
	//get these numbers in an aray
	var b []int
	for i := 0; d+i < len(a); i++ {
		b = a[i : d+i]
		fmt.Print("\n")
		fmt.Println("b is: ", b)
		median := elementing(b)
		fmt.Println("a[d+1] and median is: ", a[d+i], median)
		if float32(a[d+i]) >= 2*median {
			notification += 1
			fmt.Print("notification is: ", notification)
		} else {
			continue
		}
	}
	fmt.Println("\n notification is:", notification)
}
