package main

import "fmt"

var table = make(map[int]int)

func fibonacci(number int) int {
	if number <= 1 {
		table[number] = number
		return number
	}

	if value, found := table[number]; found {
		return value
	}

	return fibonacci(number-1) + fibonacci(number-2)
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
