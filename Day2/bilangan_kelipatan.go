package main

import (
	"fmt"
)

func isMultipleOf7(n int) bool {
	return n%7 == 0
}

func main() {
	var number int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&number)

	if isMultipleOf7(number) {
		fmt.Println(number, "adalah kelipatan dari 7.")
	} else {
		fmt.Println(number, "bukan kelipatan dari 7.")
	}
}
