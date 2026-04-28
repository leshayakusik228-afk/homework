package main

import "fmt"

func main() {
	var n int
	var a int

	fmt.Println("выбери число №1") //("выбери первое число")
	 , err := fmt.Scanln(&n)
	fmt.Println(n)

	fmt.Println("выбери другое")
	fmt.Scanln(&a)
	fmt.Println(a)

}
