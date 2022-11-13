package main

 import ( 
 "fmt" 
 "os"
 "io/ioutil"
 )

 
 func main() { 

file, _ := os.Create("test.txt")

defer file.Close()

len, _ := file.WriteString("Welcome to the party pal")
     
fmt.Println(len) 

data, _ := ioutil.ReadFile("test.txt")

fmt.Printf("\nData: %s", data)
}
