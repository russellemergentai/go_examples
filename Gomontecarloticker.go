package main

import (
	"fmt"
  "math"
  "math/rand"
  "time"
)

func calculate(ch chan float64, i int) {
  fmt.Println("thread starting ", i)
  start = time.Now()

  acc := 0
  n := 100000

  for i:=0; i<n; i++ {
    var x = rand.Float64()
    var y = rand.Float64()
    xx := Math.Pow(x, 2)
    yy := Math.Pow(y, 2)
    r := xx + yy
    if r<1.0 {
      acc++
    }

    if math.Mod(float64(i), 1000) == 0 {
      output := float64(acc)*4 / float64(i)
      ch <- output 
    }
  }
}

func main() {
	chans := []chan float64{}

  for i:=0; i<20; i++ {
    chans = append(chans, make(chan float64))
  }

  for i := range chans {
    go calculate (chans[i],i)
  }

  for range time.Tick(1000*time.Millisecond) {
    val := 0.0
    for i := range chans {
      val += <-chans[i]
    }
          fmt.Println(val/float64(len(chans)))
  }
  
}
