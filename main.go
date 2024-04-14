package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	height float32
}

type Move interface {
	Travel()
	Train()
}

func (p Person) Travel() {
	fmt.Printf("bro i'm flying fnnalyyyy my name is %s\n", p.Name)
}

func (t Person) Train() {
	fmt.Printf("brother i'm in train %s\n", t.Name)
}

func main() {
	// it is gonna be slice bcz her e2 information brother
	person := []Person{
		{Name: "Dilshod", Age: 18, height: 1.85},
		{Name: "Salombek", Age: 22, height: 195.5},
	}

	for _, n := range person {
		n.Travel()
	}

	for _, t := range person {
		t.Train()
	}
}
