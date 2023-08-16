package main

import "fmt"

func main() {
	var angka int

	fmt.Println("---MENENTUKAN BILANGAN GANJIL ATAU GENAP---")
	fmt.Print("Masukkan angka = ")
	fmt.Scan(&angka)

	if angka%2 == 0 {
		fmt.Println(angka, "= genap")
	} else {
		fmt.Println(angka, "= ganjil")
	}

}
