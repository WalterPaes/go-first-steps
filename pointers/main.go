package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{
		"Walter",
	}

	// ReferĂȘncia ao p
	p1 := &p
	p1.name = "Junior"

	// Sem referĂȘncia
	p2 := p
	p2.name = "Alberto"

	fmt.Println(p.name, p1.name, p2.name, p.name)
}
