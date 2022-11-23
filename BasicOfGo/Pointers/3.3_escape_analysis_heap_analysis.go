package main               // RETURNING POINTERS
import "fmt"               // sharing up typically escapes to the heap
func main() {
	n:= heapAnalysis()
    fmt.Println(*n/2)
}
//go:noinline
func heapAnalysis() *int {
    data := 42
    return &data
}