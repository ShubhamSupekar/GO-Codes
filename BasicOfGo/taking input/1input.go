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
var head *node = new(node)
var last *node
func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a number: ")
		myname,_ := reader.ReadString('\n')
		LL(myname)		
		for head != nil{
			fmt.Println(head.data,head.next)
		}
	}
}
func LL(myname string){
	n := node{data: myname,}
	for{
		if head==nil {
			head = &n
			last = head
			continue
		}
		last.next = &n
		last = last.next
		}
}

