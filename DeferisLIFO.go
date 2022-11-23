package main
import "fmt"

func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}

// lifo
func main() {
	b() // 3210
}
