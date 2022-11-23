package main
import "fmt"
func main(){
	arr :=[3]int{1,2,3}                                // numbers assingning in array
	arr[1]=100                                         // at index 1 in array we change value to 100
	fmt.Println("Before adding number array is",arr) 
	sum:=0                         
	for i:=0;i<len(arr);i++{
		sum += arr[i]                                   // sum of the numbers loop
	}
	fmt.Println("After adding numbers total is",sum)
	fmt.Println(len(arr))
	fmt.Println(arr[1])
}