package main

import "fmt"

func Mapping(slice []string) map[string]int {
	var result = make(map[string]int)
	var count int = 1

	for i, item1 := range slice {
		for j, item2 := range slice {
			if i == j {
				continue
			} else if item1 == item2 {
				count++
			}
		}
		result[item1] = count
		count = 1
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         // map[]
}
