package main

import "fmt"

type Person struct {
	Name string
	age  int
}

type Address struct {

}

func (p Person) getNome() string {
	return p.Name
}

func main() {
	person := Person{
		"Walter",
		26,
	}

	fmt.Printf("Name: %s | Age: %d", person.getNome(), person.age)}
