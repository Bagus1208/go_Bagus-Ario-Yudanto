package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var result []int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			if arr[i]+arr[j] == target {
				result = append(result, i, j)
				break
			}
		}
		if result != nil {
			break
		}
	}

	return result
}

//cara lain
func PairSum2(arr []int, target int) []int {
	m := make(map[int]int)

	for i, num := range arr {
		complement := target - num
		fmt.Println("nilai complemen :", complement)
		if index, found := m[complement]; found {
			fmt.Println("nilai index :", index)
			return []int{index, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	// Test cases
	fmt.Println(PairSum2([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum2([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum2([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum2([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum2([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
