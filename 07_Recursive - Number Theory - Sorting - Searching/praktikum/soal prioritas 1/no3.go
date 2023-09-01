package main

import "fmt"

func primeX(number int) int {
	var count int = 0
	var angka int = 1
	var i int

	for count < number {
		angka++

		for i = 2; i <= angka; i++ {
			if angka%i == 0 {
				break
			}
		}

		if i == angka {
			count += 1
		}
	}

	return angka
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
