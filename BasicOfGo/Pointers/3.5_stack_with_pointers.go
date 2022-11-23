package main              //sharing down typically stays on the stack
import "fmt"
func main(){
	n:=4
	inc(&n)
	fmt.Println(n)
}
func inc(a*int){
	*a++
}