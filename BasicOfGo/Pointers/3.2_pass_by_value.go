package main                          // Sharing Data
                                      // in this program we have to change the value of "hello" to "changed!" for this
									  // we tried both ways using pointer or not usint it 
import	"fmt"

func changevalue(str *string){          // this function is using pointers 
	*str = "Changed!"
}
func changevalue2(str string){           // this function is not using pointers
	str = "Changes!"
}
func main() { 
	tochange := "Hello"
	print(tochange)                  // before applying any function
	print("\n",&tochange)
	changevalue2(tochange)           //this gives the result for not using pointers
	fmt.Println("\nNot using pointers we get",tochange)
	
	changevalue(&tochange)
	fmt.Println("Using pointers we get",tochange)
	print(&tochange)

	

	
}


