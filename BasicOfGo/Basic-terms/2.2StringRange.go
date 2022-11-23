package main
import "fmt"

func main(){
	a := "abcdefghijklmnopqrstuvwxyz"
	fmt.Print(len(a))
	for d,c:=range a{
		fmt.Println("\n",d,c)
	}
}