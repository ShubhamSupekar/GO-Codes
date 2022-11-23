package main
import "fmt"
import "strings"
func main(){
	name := "hi my name is Shubham"
	numbers:= "123014562447899510"
	splitName1 := strings.Split(name,",")
	splitName2 := strings.Split(name,"!")
	splitName3 := strings.Split(name,"")
	splitNumbers := strings.Split(numbers,"")
	fmt.Println(splitName1,splitName2,splitName3)
	fmt.Println(splitNumbers)
}