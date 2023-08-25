package main

import "fmt"

func getMinMax(numbers ...*int) (min, max int) {
	min = 0
	max = 0

	for i := 1; i < len(numbers); i++ {
		if *numbers[min] > *numbers[i] {
			min = i
		}
		if *numbers[max] < *numbers[i] {
			max = i
		}
	}

	min = *numbers[min]
	max = *numbers[max]

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Print("Input angka ke-1 : ")
	fmt.Scan(&a1)
	fmt.Print("Input angka ke-2 : ")
	fmt.Scan(&a2)
	fmt.Print("Input angka ke-3 : ")
	fmt.Scan(&a3)
	fmt.Print("Input angka ke-4 : ")
	fmt.Scan(&a4)
	fmt.Print("Input angka ke-5 : ")
	fmt.Scan(&a5)
	fmt.Print("Input angka ke-6 : ")
	fmt.Scan(&a6)
	fmt.Println()

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min : ", min)
	fmt.Println("Nilai Max : ", max)
}
