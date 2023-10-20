package main

import (
	"fmt"
)

func Mapping(slice []string) map[string]int {
	result := make(map[string]int)
	for _, str := range slice {
		result[str]++
	}
	return result
}

func main() {
	s := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	mapped := Mapping(s)
	fmt.Println(mapped) // Output: map[apple:3 banana:2 orange:1]
}
