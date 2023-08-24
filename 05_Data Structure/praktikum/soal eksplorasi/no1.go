package main

import (
	"fmt"
)

func BuatMatrik(baris, kolom int) [][]int {
	var matrik = make([][]int, baris)
	var angka int

	for i := 0; i < len(matrik); i++ {
		matrik[i] = make([]int, kolom)
	}

	for i := 0; i < baris; i++ {
		for j := 0; j < kolom; j++ {
			fmt.Print("Masukkan nilai matrik baris ", i+1, " kolom ", j+1, " : ")
			fmt.Scanln(&angka)
			matrik[i][j] = angka
		}
	}

	for i := 0; i < baris; i++ {
		fmt.Println(matrik[i])
	}
	return matrik
}

func SelisihDiagonal(matrik [][]int) int {
	var hasil1, hasil2 int

	for i := 0; i < len(matrik); i++ {
		hasil1 += matrik[i][i]
		hasil2 += matrik[i][len(matrik)-1-i]
	}
	if hasil1 < hasil2 {
		return hasil2 - hasil1
	}

	return hasil1 - hasil2
}

func main() {
	// var matrik = [][]int{
	// 	{5, 2, 1},
	// 	{4, 5, 6},
	// 	{9, 8, 10},
	// }
	var matrik int

	matrik = SelisihDiagonal(BuatMatrik(3, 3))
	fmt.Println("Selisih diagonal matrik adalah", matrik)
}
