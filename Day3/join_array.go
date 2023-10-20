package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	mergedArray := make([]string, 0, len(arrayA)+len(arrayB)) // preallocate the underlying array for optimization
	mergedArray = append(mergedArray, arrayA...)
	mergedArray = append(mergedArray, arrayB...)
	return mergedArray
}

func main() {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"d", "e", "f"}
	merged := ArrayMerge(arr1, arr2)
	fmt.Println(merged)
}
