package main
import(
	"os"
	"fmt"
	"bufio"
	"strconv"

)
func main(){
reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	myname,_ := reader.ReadString('\n')
	mynum,_:= strconv.ParseInt(myname,1,2) 
	fmt.Println(myname)
	fmt.Println(mynum)
}