package main          // RETURNING POINTERS
import "fmt"          // sharing up typically escapes to the heap

func main(){
	n:= answer()
	fmt.Println(*n/2)
}
func answer()*int{
	data:=42
	return &data
}