package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&number)

	if isPrime(number) {
		fmt.Println(number, "adalah bilangan prima.")
	} else {
		fmt.Println(number, "bukan bilangan prima.")
	}
}
