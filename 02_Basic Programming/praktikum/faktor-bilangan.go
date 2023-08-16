package main

import "fmt"

func main() {
	var angka int

	fmt.Println("---FAKTOR BILANGAN---")
	fmt.Print("Masukkan angka = ")
	fmt.Scan(&angka)
	fmt.Print("Faktor bilangan dari ", angka, " = ")

	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
