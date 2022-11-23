package main

import "fmt"

func main(){
	 arr := []int{1,1,0,-1,-1}
	 var pos,nev,zer float64


	 for _, num := range arr{
		if num>0 {
			pos++
		}else if num<0 {
			nev++
		}else {
			zer++
		} 
	 }
	  var result float64 
	   n := float64(len(arr))
	 fmt.Println(pos,zer,nev)
	 result = pos/n
	 fmt.Println("")
	 fmt.Printf("%.6f",result)
	 fmt.Println("")

	 fmt.Println(result)
	 fmt.Println(result)


}