package main

import "fmt"

func main() {
	var kata, result string

	fmt.Println("---MENGECEK POLINDROME---")
	fmt.Print("Masukkan kata = ")
	fmt.Scan(&kata)

	for i := len(kata) - 1; i >= 0; i-- {
		result += string(kata[i])
	}
	fmt.Println("captured =", result)

	if result == kata {
		fmt.Println("Kata tersebut merupakan Polindrome")
	} else {
		fmt.Println("Kata tersebut bukan Polindrome")
	}
}
