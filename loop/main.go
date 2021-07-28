package main

import "fmt"

func main() {
	/*var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)*/

	/*for {
		fmt.Println("Hello")
	}*/

	numbers := []int{1, 22, 54, 15, 66, 97}
	for i, v := range numbers {
		fmt.Printf("Index: %d | Value: %d\n", i, v)
	}
}
