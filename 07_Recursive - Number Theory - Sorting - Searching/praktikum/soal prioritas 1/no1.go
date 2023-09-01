package main

import "fmt"

func fibonacci(angka int) int {
	if angka <= 1 {
		return angka
	}

	return fibonacci(angka-1) + fibonacci(angka-2)
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
