package main

import "fmt"

func main() {
	lang := "es"

	switch lang {
	case "pt":
		fmt.Println("Olá")
	case "en":
		fmt.Println("Hello")
	case "es":
		fmt.Println("Hola")
	default:
		fmt.Println("...")
	}
}
