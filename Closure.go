
/*
A closure is like a little function "package" that contains 
both a function and the data it needs to work with. 
It "closes over" or captures the variables from the surrounding scope, 
allowing the function to access and manipulate those variables even after 
the surrounding scope has finished executing. 
In simple terms, it's a function that remembers and uses variables from the place where it was created.
*/

package main

import "fmt"

func main() {
    // Define a function that returns a closure
   // Ie 'func() int'
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }
