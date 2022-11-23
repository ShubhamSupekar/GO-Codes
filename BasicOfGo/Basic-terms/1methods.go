package main

import "fmt"

type student struct {
	name   string
	grades []int
	age    int
}

func (s *student) setAge(age int) {
	s.age = age
}

func (s student) getAvgGrade() float32 {
	sum := 0
	for _,v := range s.grades {
		sum += v
	}
	return float32(sum)/ float32(len(s.grades))
}
func (s *student)getMax() int{
	currMax := 0
	for _,v := range s.grades{
		if currMax <v{
			currMax = v 
		}
	}
	return currMax
}

func main() {
	s1 := student{"tim", []int{70, 40, 50, 60}, 19}
	s2 := student{"tim", []int{0, 4, 0, 6}, 9}
	fmt.Println(s1)
	s1.setAge(7)
	fmt.Println(s1)
	avg:= s1.getAvgGrade()
	avg1 := s2.getAvgGrade()
	max1 := s1.getMax() 
	max2:= 	s2.getMax()
	fmt.Println(avg,avg1)
	fmt.Println(max1,max2)
}
