package main

import "fmt"

func main() {
	var input, output []int
	var number int
	for {
		fmt.Println("Enter the number(or press q to END): ")
		_, err := fmt.Scan(&number)
		if err == nil {
			input = append(input, number)
			if len(output) == 0 {
				output = append(output, recursion(number))
				fmt.Println(recursion(number))
				continue
			}
			result := check(number, output, input)
			output = append(output, result)
			fmt.Println(result)
			continue
		} else {
			fmt.Println("q to exit check")
			for j := 0; j < len(output); j++ {
				fmt.Printf("\n%d factorial is %d", input[j], output[j])
			}
			fmt.Println("\n", input, output)
			break
		}
	}
}
func check(a int, b, f []int) (temp2 int) {
	check := false
	for i := 0; i < len(b); i++ {
		temp1 := a - 1
		if temp1 == f[i] {
			fmt.Println("1 less value found")
			temp2 = b[i] * a
			check = true
		}
		temp3 := a + 1
		if temp3 == f[i] {
			fmt.Println("1 greater value found")
			temp2 = b[i] / temp3
			check = true
		}
	}
	if check == false {
		fmt.Println("value not found")
		temp2 = recursion(a)
		return
	}
	return
}
func recursion(d int) int {
	if d == 1 {
		return 1
	}
	return recursion(d-1) * d
}
