package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	counts := make(map[rune]int)
	for _, digit := range angka {
		counts[digit]++
	}

	var result []int
	for _, digit := range angka {
		if counts[digit] == 1 {
			result = append(result, int(digit-'0')) // Convert rune to integer
		}
	}
	return result
}

func main() {
	num := "12345566778"
	uniqueDigits := munculSekali(num)
	fmt.Println(uniqueDigits) // Output: [1 2 3]
}
