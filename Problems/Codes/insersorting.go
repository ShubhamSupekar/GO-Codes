package main
import "fmt"
func main() {
// index starts from 0
var scores = []int{90, 70, 50, 80, 60, 85}
var length = len(scores)
sort(scores, length)
for i := 0; i < length; i++ {
fmt.Printf("%d,", scores[i])
}
}
func sort(arrays []int, length int) {
for i := 0; i < length; i++ {
var insertElement = arrays[i] //Take unsorted new elements
var insertPosition = i //Inserted position
for j := insertPosition - 1; j >= 0; j-- {
//If the new element is smaller than the sorted element, it is
//shifted to the right
if insertElement < arrays[j] {
arrays[j+1] = arrays[j]
insertPosition--
}
}
arrays[insertPosition] = insertElement //Insert the new element
}
}