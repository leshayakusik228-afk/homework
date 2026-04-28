package main

import "fmt"

func sum(a int64, b int64) int64 {
	return a + b

}
func main() {
	var a int64 = 10
	var b int64 = 12

	result := sum(a, b)
	fmt.Println(result)
}
