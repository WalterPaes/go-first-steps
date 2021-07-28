package main

import "fmt"

func main() {
	// Maps são coleções de chave e valor
	person := make(map[string]string)
	person["name"] = "Walter"
	person["surname"] = "Junior"

	fmt.Println(person)

	// Declaração Literal
	students := map[int]string{
		1: "João",
		2: "Maria",
	}

	fmt.Println(students)

	/*var fruits map[int]string
	fruits[1] = "maçã"
	fruits[2] = "banana"

	fmt.Println(fruits)*/
}
