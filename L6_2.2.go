
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
	sliceTree := append(firstSlice, twoSlice...)
	//sliceTree[4] = 32
	//sliceTree = append(sliceTree, 12, 45, 44)

	fmt.Println(firstSlice)
	fmt.Println(twoSlice)
	fmt.Println(sliceTree)
	//fmt.Println("длина=", len(sliceTree), "вместимость", cap(sliceTree))
}
