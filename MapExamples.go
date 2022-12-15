package main

import (
	"fmt"
	// myPackage "main/myPackage"
)

func main() {
	// definition
	var names = make(map[int]string)
	names[0] = "Fred"
	names[1] = "Jane"
  names[2] = "Chad"

	fmt.Println(names)
	
  // existence 
	fred, exists := names[0]
    if(exists) {
        fmt.Printf("%s exists \n", fred)  // prints "Freddy exists"
    }   
		
// delete
delete (names,1)
  
  // iterate
  for _, name := range names {
        fmt.Println(name)
    }

}
