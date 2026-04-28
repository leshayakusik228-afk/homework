package main

import "fmt"

func main() {
	var x int
	x = 0

	if x > 10 {
		fmt.Println("результат большой")
	} else if x >= 1 {
		fmt.Println("результат средний")

	} else {
		fmt.Println("результат маленький или 0")
	}
}
