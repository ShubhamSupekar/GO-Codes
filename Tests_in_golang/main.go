package main

import "fmt" 
func addition(a int)(int){
	result := a+2
	return result
}

func main(){
	c:= addition(2)
	fmt.Println(c)
}