package main
//append an element in any position in LL
import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func main() {
	var address *node
	ip := []int{1, 2, 3, 4, 5}
	var position int = 4
	var data int = 60
	var counter int = 0
	for _, val := range ip {
		n := node{
			data: val,
		}
		if address == nil {
			address = &n
			continue
		}
		last := address
		for last.next != nil {
			last = last.next
		}
		last.next = &n
	}
	toprint := address
	for toprint != nil {
		fmt.Println("Before change ll is: ", toprint.data, toprint.next)
		toprint = toprint.next
	}
	past := address

	if position >= len(ip) {
		fmt.Println("exceed the limit")
	} else {
		for counter = 0; ; {
			if counter == position {
				past.data = data
				break
			}
			past = past.next
			counter++
			continue
		}
	}
	for address != nil {
		fmt.Println("after change ll is: ", address.data, address.next)
		address = address.next
	}
}
