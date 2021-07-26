package main

import "errors"

func get() error {
	return errors.New("an error")
}

func main() {
	err := get()
	if err != nil {
		panic(err)
	}
}
