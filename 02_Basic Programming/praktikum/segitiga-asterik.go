package main

import "fmt"

func main() {
	var baris int

	fmt.Println("---SEGITIGA ASTERIK---")
	fmt.Print("Berapa baris yang ingin dicetak = ")
	fmt.Scanln(&baris)

	for i := 1; i <= baris; i++ {
		for j := baris; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k < 2*i; k++ {
			if k%2 != 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
