package main
import "fmt"
func main() {
    // Write your code here
	a:= []int{1, 2 ,3, 4, 5}
	d := 5
    for counter := 1;counter <d;counter++{
        b:= a[0]
        a = a[1: ]
        a= append(a,b)
    }
	fmt.Println(a)
}