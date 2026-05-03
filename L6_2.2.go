package main

import "fmt"

func doubleElements(slice []int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = v * 2

	}
	return result
}

func main() {
	firstSlice := []int{1, 2, 3, 4, 5}
	twoSlice := doubleElements(firstSlice)
	fmt.Println(firstSlice)
	fmt.Println(twoSlice)
}
