package main

import "fmt"

func PascalTriangle(row int) [][]int {
	var result = make([][]int, row)

	for i := 1; i <= 5; i++ {
		result[i-1] = make([]int, i)
		number := 1

		for j := 1; j <= i; j++ {
			result[i-1][j-1] = number
			number = number * (i - j) / j
		}
	}

	return result
}

func main() {
	var row int

	fmt.Print("Input baris : ")
	fmt.Scan(&row)
	fmt.Println(PascalTriangle(row))
}
