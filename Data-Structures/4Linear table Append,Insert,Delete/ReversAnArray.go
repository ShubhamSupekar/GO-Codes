package main
import(
	"fmt"
)
func main(){
	a:=[]int{20,30,40,60,80,90}
	length := len(a)
	for i:=0;i<len(a)/2;i++{

		temp:= a[i]
		a[i] = a[length-i-1]
		a[length-i-1] = temp
	}
	fmt.Println(a)

}