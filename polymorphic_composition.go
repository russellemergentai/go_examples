package main

import (
	"fmt"
)

// setup interface for polymorphic behavior
type Worker interface {
	Work()
}

// create person, developer, manager with composition *instead of inheritance*
type Person struct {
	Name string
}

type Developer struct {
	Person         // this is composition, not inheritance
	Specialization string
}

type Manager struct {
	Person
	Team string
}

// struct receivers with the composition structs to define polymorphic behaviour, based in the interface
func (d Developer) Work() {
	fmt.Println("%s working", d.Name)
}

func (m Manager) Work() {
	fmt.Println("%s slacking", m.Name)
}

// define a polymorphic array and run a polymorphic loop on the interface
func main() {
	workers := []Worker{
		Developer{Person: Person{Name: "bob"}, Specialization: "golang"},
		Manager{Person: Person{Name: "git"}, Team: "slacking"},
	}

	for i, worker := range workers {
		fmt.Println(i)
		worker.Work()
	}

}
