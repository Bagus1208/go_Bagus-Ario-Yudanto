package main

import (
	"fmt"
	"strconv"
)

func ToBinary(n int) []string {
	var result = make([]string, n+1)
	var value string

	for i := 0; i <= n; i++ {
		value = strconv.FormatInt(int64(i), 2)
		result[i] = value
	}

	return result
}

func main() {
	fmt.Println(ToBinary(2))
	fmt.Println(ToBinary(5))
	fmt.Println(ToBinary(10))
}
