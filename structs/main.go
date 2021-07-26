package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) getNome() string {
	return p.Name
}

func main() {
	person := Person{
		"Walter",
		26,
	}

	fmt.Println("Name:" + person.getNome())
}
