package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	var (
		result string
		temp   byte
		key    = byte(offset % 26)
	)

	for i := 0; i < len(input); i++ {
		if input[i] == 32 {
			result += string(input[i])
			continue
		} else if input[i] <= 90 {
			temp = input[i] - 65
			temp = (temp + key) % 26
			temp = temp + 65
		} else {
			temp = input[i] - 97
			temp = (temp + key) % 26
			temp = temp + 97
		}

		result += string(temp)
	}

	return result
}

func main() {

	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // def
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
	fmt.Println(caesar(10, "Bagus Ario Yudanto"))           // Lkqec Kbsy Ienkxdy
}
