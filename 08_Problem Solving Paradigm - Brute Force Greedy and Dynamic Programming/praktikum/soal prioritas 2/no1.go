package main

import (
	"fmt"
)

func Absolute(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

func Frog(jumps []int) int {
	var result int
	var rock int = 0

	for i := 0; i < len(jumps)/2; i++ {
		if rock == len(jumps)-2 {
			result += Absolute(jumps[rock+1], jumps[rock])
		} else if Absolute(jumps[rock+1], jumps[rock]) > Absolute(jumps[rock+2], jumps[rock]) {
			result += Absolute(jumps[rock+2], jumps[rock])
			rock += 2
		} else {
			result += Absolute(jumps[rock+1], jumps[rock])
			rock++
		}
	}

	return result
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))                  // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))          // 60
	fmt.Println(Frog([]int{70, 50, 60, 40, 90, 30, 100, 40})) // 50
	fmt.Println(Frog([]int{70, 50, 60}))                      // 10

}
