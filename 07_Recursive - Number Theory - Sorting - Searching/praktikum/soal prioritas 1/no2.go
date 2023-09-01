package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	var tmp = make(map[string]int)

	for _, item := range items {
		current := item

		if _, found := tmp[current]; found {
			tmp[current]++
		} else {
			tmp[current] = 1
		}
	}

	var result = make([]pair, len(tmp))
	var index int = 0

	for key, item := range tmp {
		result[index].name = key
		result[index].count = item
		index++
	}

	result = bubblesort(result)

	return result
}

func bubblesort(element []pair) []pair {
	for i := 0; i < len(element)-1; i++ {
		for j := 0; j < len(element)-1-i; j++ {
			if element[j].count > element[j+1].count {
				element[j].count, element[j+1].count = element[j+1].count, element[j].count
			}
		}
	}
	return element
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
}
