//Add a score 75 to the end of the one-dimensional array scores.

package main
import(
	"fmt"
) //append 75 element in the array
func main(){
	a := []int{90,70,50,80,60,85}
	length := len(a)
	temp := make([]int,length+1)
	for i:=0;i<length;i++{
		temp[i] = a[i]
	}
	temp[length]=75
	a = temp
	fmt.Println(a)
}