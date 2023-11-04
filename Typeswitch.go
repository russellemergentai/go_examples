package main

import "fmt"

/*
No Methods: An empty interface doesn't require the values it holds to implement any methods or have any specific behavior.
Universal Container: It serves as a universal container for values of different types. You can use an empty interface to store values of various data types in a single variable.
Type Assertion: To work with the values stored in an empty interface, you typically use type assertion or type switching to determine the actual type of the value and perform appropriate operations.

*/


func main() {
    // Declare an empty interface variable 'val'
    var val interface{}

    // Assign an integer (42) to the empty interface 'val'
    val = 42

    // Use a type switch to check the actual type of 'val'
    switch val.(type) {
    case int:
        fmt.Println("It's an integer!")
    case string:
        fmt.Println("It's a string!")
    default:
        fmt.Println("It's a different type.")
    }
}


