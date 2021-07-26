package main

import "fmt"

func main() {
	// Inicializando variáveis com valores zerados
	var name string
	name = "Walter"

	// Variáveis com inferência do tipo
	age := 26

	// Variáveis com o tipo explícito
	var email string = "walter@email.com"

	fmt.Printf("Name: %s \n", name)
	fmt.Printf("Age: %d \n", age)
	fmt.Printf("Email: %s \n", email)
}
