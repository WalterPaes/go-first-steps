package main

import (
	"fmt"
)

func main() {
	// Arrays possuem tamanho fixo e um único tipo
	var people [4]string // nil
	people[0] = "Tuba"
	people[1] = "Ulisses"
	people[2] = "Walter"
	people[3] = "Josiano"

	// Declaração literal
	people2 := [3]string{"Fabrício", "Richard", "Israel"}

	fmt.Println(people, people2)
}
