package main

import "fmt"

func main() {
	// Slices são "arrays" dinâmicos
	var animes []string
	fmt.Println(animes)

	animes = append(animes, "Naruto")
	fmt.Println(animes)

	animes = append(animes, "DBZ")
	fmt.Println(animes)

	// Declaração literal
	movies := []string{"Spider-man", "Iron man", "Thor"}
	fmt.Println(movies)
}
