package main

import "fmt"

type person interface {
	talk()
}

type walter struct{}

func (w walter) talk() {
	fmt.Println("Hello")
}

func main() {
	var p person

	p = walter{}

	p.talk()
}
