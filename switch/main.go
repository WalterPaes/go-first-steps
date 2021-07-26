package main

import "fmt"

func main() {
	lang := "pt"

	switch lang {
	case "pt":
		fmt.Println("Olá")
	case "en":
		fmt.Println("Hello")
	case "es":
		fmt.Println("HOla")
	default:
		fmt.Println("...")
	}
}
