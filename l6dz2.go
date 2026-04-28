package main

import "fmt"

func devide(a int, b int) (int, int) {
	quo := a / b // деление
	rem := a % b // остаток

	return quo, rem
}
func main() {
	var a int = 11
	var b int = 2

	q, r := devide(a, b)
	fmt.Println(q)
	fmt.Println(r)

}
