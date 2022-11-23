package main             // THIS PROGRAM IS IN FINITE LOOP 
import "fmt"
const size = 10
func main() {
	s := "Hello"
	stackcopy(&s,0)
}
func stackcopy(s *string, c int) {
	fmt.Println(c,s,*s)
	c++
	if c ==10{
		
		return 
	}
	stackcopy(s,c)
}
