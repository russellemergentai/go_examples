package main
import "fmt"

func main() {
	fmt.Printf("Hello, " + GetInstanceA().str)
}

type A struct {
 str string
}

var singleton *A
func init() {
 //initialize static instance on load
 singleton = &A{str :"abc"} 
}
//GetInstanceA - get singleton instance pre-initialized
func GetInstanceA() *A {
 return singleton
}
