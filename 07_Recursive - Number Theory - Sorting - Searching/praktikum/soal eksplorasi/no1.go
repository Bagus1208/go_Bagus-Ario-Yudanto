package main

import (
	"fmt"
)

func MaxSequence(arr []int) int {
	var result int
	var max int

	for i := 0; i < len(arr); i++ {
		result += arr[i]

		if result < arr[i] {
			result = arr[i]
		}

		if max < result {
			max = result
		}
	}

	return max
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))      // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))        // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))        // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))        // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))         // 12
	fmt.Println(MaxSequence([]int{7, 9, -20, 18, -30, 54, 2, -10, 70})) //116
}
