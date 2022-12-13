package main

import (
	"fmt"
	// myPackage "main/myPackage"
)

type Human interface {
  hobby() string
}

type Man struct {
  name string
}

type Woman struct {
  name string
}

// implement the interface for each struct
func (m Man) hobby () string {
 return m.name
}

// note each struct is an argument of a function implementing the interface method 
// making man, woman <= human 
func (m Woman) hobby () string {
  return m.name
}
      
func main() {
  m := new (Man)
  m.name = "chad"
  w := new (Woman)
  w.name = "karen"

  fmt.Println(m.hobby())
  fmt.Println(w.hobby())

  // do polymorphism via interface array

  people := []Human{m, w}

  for _, item := range people {
    fmt.Printf("%s \n", item.hobby())
  }

}
