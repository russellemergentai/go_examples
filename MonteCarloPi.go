package main

import (
 "fmt"
 "math"
 "math/rand"
 ) 
   
 
 func main() {
   acc := 0
   n := 10000000
   
   for i := 0; i < n; i++ { 	
       var x = rand.Float64()
       var y = rand.Float64()
       xx := math.Pow(x,2)
       yy := math.Pow(y,2)
       r := (xx + yy)
       if (r < 1.0) {acc++}
   }
   
   output := float64(acc) * 4 / float64(n)
   
   fmt.Println(output)
 }
