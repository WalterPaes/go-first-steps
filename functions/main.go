package main

import "fmt"

// Definindo uma função com parâmetro e retorno
func add(x int, y int) int {
	return x + y
}

// Quando o tipo dos parâmetros são iguais, podemos omitir o tipo do primeiro parâmetro
func concat(a, b string) string {
	return a + b
}

// Retornando múltiplos valores
func example() (string, string) {
	return "foo", "bar"
}

// Função sem parâmetro e sem retorno
func main() {
	result := add(1, 2)
	fmt.Println(result)

	word := concat("Hello", "World")
	fmt.Println(word)

	_, word2 := example()
	fmt.Println(word2)
}
