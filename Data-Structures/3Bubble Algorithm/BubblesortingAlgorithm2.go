package main

import "fmt"
//sorting using functions
func main() {
	score := []int{20, 60, 90, 40, 00}
	length := len(score)
	sorts(score, length)
	for k := 0; k < length; k++ {
		fmt.Printf("%d,", score[k])
		
	}
	
}
func sorts(score []int, length int) {
	for i := 1; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if score[j] > score[j+1] {
				//swap
				temp := score[j]
				score[j] = score[j+1]
				score[j+1] = temp
			}
		}
	}
	fmt.Println("max number is:",score[length-1])
}
