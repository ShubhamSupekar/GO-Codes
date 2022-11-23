package main             // THIS PROGRAM IS IN INFINITE LOOP 
import "fmt"
const size = 10
func main() {
	s := "Hello"
	stackcopy(&s,0,[size]int{})
}
func stackcopy(s *string, c int, a [size]int) {
	fmt.Println(c,s,*s)
	if c ==10{
		
		return 
	}
	stackcopy(s,c,a)
}
