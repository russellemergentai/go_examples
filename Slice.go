package main

import (
  "fmt"
)

func main() {
    // Initializing a slice from an array
    array := [5]int{1, 2, 3, 4, 5}
    slice := array[1:4]

    // Displaying the original array and slice
    fmt.Println("Original array:", array)
    fmt.Println("Initial slice:", slice)

    // Mutability
    slice[0] = 99

    // Displaying the array and slice after mutation
    fmt.Println("Array after mutation:", array)
    fmt.Println("Slice after mutation:", slice)

    // Resizing the slice
    resizedSlice := array[:2]

    // Displaying the array and resized slice
    fmt.Println("Array after resizing:", array)
    fmt.Println("Resized slice:", resizedSlice)
}
