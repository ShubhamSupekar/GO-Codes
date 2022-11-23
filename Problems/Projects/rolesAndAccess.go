package main

//we can give access like this
import (
	"fmt"
)

const (
	isAdmin = 1 << iota
	isRoot
	isGuest
	isPeoples
)

func main() {
	var access byte = isAdmin | isRoot
	fmt.Printf("all deta of access is %b\n", access)
	fmt.Printf("is Admin have access %v\n", isAdmin&access == isAdmin)
	fmt.Printf("is Guest have access %v", isGuest&access == isGuest)

}
