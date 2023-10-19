package main

import (
	"fmt"
)

// Fungsi untuk menghitung luas trapesium
func trapezoidArea(a float64, b float64, h float64) float64 {
	return 0.5 * (a + b) * h
}

func main() {
	var a, b, h float64

	fmt.Print("Masukkan panjang sisi atas trapesium (a): ")
	fmt.Scan(&a)

	fmt.Print("Masukkan panjang sisi bawah trapesium (b): ")
	fmt.Scan(&b)

	fmt.Print("Masukkan tinggi trapesium (h): ")
	fmt.Scan(&h)

	area := trapezoidArea(a, b, h)
	fmt.Printf("Luas trapesium dengan sisi atas %.2f, sisi bawah %.2f, dan tinggi %.2f adalah: %.2f\n", a, b, h, area)
}
