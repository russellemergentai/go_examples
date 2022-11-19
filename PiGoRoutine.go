package main

import (
	"fmt"
	"math"
	"math/rand"
)

func calculate(ch chan float64) {
	acc := 0
	n := 10000000

	for i := 0; i < n; i++ {
		var x = rand.Float64()
		var y = rand.Float64()
		xx := math.Pow(x, 2)
		yy := math.Pow(y, 2)
		r := (xx + yy)
		if r < 1.0 {
			acc++
		}
	}

	output := float64(acc) * 4 / float64(n)
	ch <- output
}

func main() {
	ch1 := make(chan float64)
	go calculate(ch1)
// Blocks until data available 
	val := <-ch1
	fmt.Println(val)
}
