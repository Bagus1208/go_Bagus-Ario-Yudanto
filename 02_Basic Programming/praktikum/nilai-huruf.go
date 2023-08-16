package main

import "fmt"

func main() {
	var nilai int

	fmt.Println("---MENENTUKAN NILAI HURUF---")
	fmt.Print("Masukkan nilai Anda = ")
	fmt.Scan(&nilai)

	switch {
	case nilai > 100 || nilai < 0:
		fmt.Println("Nilai invalid")
	case nilai >= 80:
		fmt.Println("Anda mendapat nilai huruf A")
	case nilai >= 65:
		fmt.Println("Anda mendapat nilai huruf B")
	case nilai >= 50:
		fmt.Println("Anda mendapat nilai huruf C")
	case nilai >= 35:
		fmt.Println("Anda mendapat nilai huruf D")
	case nilai >= 0:
		fmt.Println("Anda mendapat nilai huruf E")
	}
}
