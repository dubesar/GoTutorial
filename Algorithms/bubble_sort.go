package main

import "fmt"
func BubbleSort(arr[] int)[]int {
   for i:=0; i< len(arr)-1; i++ {
      for j:=0; j < len(arr)-i-1; j++ {
         if (arr[j] > arr[j+1]) {
            arr[j], arr[j+1] = arr[j+1], arr[j]
         }
      }
   }
   return arr
}
func main() {
   arr:= []int{11, 14, 3, 8, 18, 17, 43};
   fmt.Println(BubbleSort(arr))
}