package main

type person struct {
	name string
}

func main() {
	p := person{
		"Walter",
	}

	// Referência ao p
	p1 := &p
	p1.name = "Junior"

	// Sem referência
	p2 := p
	p2.name = "Alberto"

	println(p.name, p1.name, p2.name, p.name)
}
