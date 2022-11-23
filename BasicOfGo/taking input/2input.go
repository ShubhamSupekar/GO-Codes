package main

import (
	"bufio"
	"fmt"
	"os"
)
type node struct{
	data string
	next *node
} 

func main(){
	var head *node 
	var last *node 
	//myname := []int{10,20,30,40,50}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a number: ")
		myname,_ := reader.ReadString('\n')
		n := node{data: myname,}
		if head==nil {
				head = &n
				last = head
				continue
		}
		last.next = &n
		last = last.next
		for head !=nil{
			fmt.Println(head.data,head.next)
			head = head.next
		}
	}
	
}




