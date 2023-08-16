package main

import "fmt"

func main() {
	var a, b, t, Luas int

	fmt.Println("---MENGHITUNG LUAS TRAPESIUM---")
	fmt.Print("Masukkan panjang sisi atas (cm) = ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan panjang sisi bawah (cm) = ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan Tinggi (cm) = ")
	fmt.Scanln(&t)

	Luas = ((a + b) * t) / 2

	fmt.Println("Luas trapesium =", Luas, "cm^2")

}
