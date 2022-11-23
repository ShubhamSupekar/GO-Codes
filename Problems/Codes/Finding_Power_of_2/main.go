package main


//unfinished

import (
	"fmt"
	"strconv"
)
func CheckNumberPowerOfTwo(a string) (int64, bool) {
   var count int64 
   var check bool
   var in int64
   for index, val := range a{
      if val == 47{
         in = int64(index)
         count =+1
      }
      if count == 2{
         check = false
         return 0, check
      }
      fmt.Println(count)
      continue
   }
   if count == 0{
      check = false
      return 0, check
   }
   check = true
   return in, check
}
func main(){
   var n = 2
   fmt.Printf("Binary of %d is: %s.\n", n, strconv.FormatInt(int64(n), 2))
   a:= strconv.FormatInt(int64(n), 2)
   flag,TF := CheckNumberPowerOfTwo(a)
   if TF == true{
      fmt.Printf("%d is the %dth power of 2",n,flag)
   }else {
      fmt.Printf("%d is not in the power of 2",n)      
   }
}