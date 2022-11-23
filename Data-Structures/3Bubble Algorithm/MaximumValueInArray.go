package main
import(
	"fmt"
)
func main(){
	a := []int{180,3,6000,122,10,10000000,20}
	max := a[0]

    for i := 1; i < len(a); i++ {

        if max < a[i] {

            max = a[i]
        }
    }

    fmt.Printf("\nMax element is : %d", max)
	
}