package main
import(
	"fmt"
)
//Initial value minIndex=0, j=1 Compare arrays[minIndex] with arrays[j]
//if arrays[minIndex] > arrays[j] then minIndex=j, j++ else j++. continue until
//the last number, arrays[minIndex] is the Min Value
func main(){
	arrays := []int{60,80,95,50,70}
	minIndex := 0
	for j:=1;j<len(arrays);j++{
		if arrays[minIndex]>arrays[j]{
			minIndex = j
			j++
		}else {
			j++
			continue
		}
	}
	fmt.Println(arrays[minIndex])
	fmt.Println(arrays)
}