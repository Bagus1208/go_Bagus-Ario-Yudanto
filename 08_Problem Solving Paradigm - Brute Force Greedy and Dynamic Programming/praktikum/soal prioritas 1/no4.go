package main

import "fmt"

func sum(a, b, c, target int) bool {
	return a+b+c == target
}
func multiple(a, b, c, target int) bool {
	return a*b*c == target
}
func power(a, b, c, target int) bool {
	return (a*a)+(b*b)+(c*c) == target
}

func SimpleEquations(a, b, c int) any {
	for i := 1; i <= a; i++ {
		for j := i; j <= a; j++ {
			for k := j; k <= a; k++ {
				if sum(i, j, k, a) && multiple(i, j, k, b) && power(i, j, k, c) {
					return []int{i, j, k}
				}
			}
		}
	}

	return "no solution"
}

func main() {
	fmt.Println(SimpleEquations(1, 2, 3))      // no solution
	fmt.Println(SimpleEquations(6, 6, 14))     // 1 2 3
	fmt.Println(SimpleEquations(15, 120, 77))  // 4 5 6
	fmt.Println(SimpleEquations(29, 400, 441)) // 4 5 20
}
