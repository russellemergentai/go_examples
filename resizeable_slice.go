// create a shape interfqce, define an array of shape structs, with an area function, 
// pass an array of shgapes to a calculate function

package main

import (
  "fmt"
)

func main() {
  // A Go Slice contains three elements: data, length, and capacity.
  s := make([]int, 0, 10)

  fmt.Println("length of elements: ", len(s))
  fmt.Println("capacity of elements: ", cap(s))

  s = []int{1,2,3,4,5}
  
  fmt.Println("new length of elements: ", len(s))
  fmt.Println("capacity of elements: ", cap(s))
  fmt.Println(s)

  // increase capacity of slice
  t := make([]int, len(s), 20)
  copy(t, s)

  fmt.Println("length of elements: ", len(t))
  fmt.Println("new capacity of elements: ", cap(t))
  fmt.Println(t)
  
}
