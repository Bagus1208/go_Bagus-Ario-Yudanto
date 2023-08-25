package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var result float64

	for _, value := range s.score {
		result += float64(value)
	}

	return result / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]

	for i, value := range s.score {
		if min > value {
			min = value
			name = s.name[i]
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]

	for i, value := range s.score {
		if max < value {
			max = value
			name = s.name[i]
		}
	}

	return max, name
}

func main() {
	var a = Student{}
	var total int

	fmt.Print("Input number of student : ")
	fmt.Scanln(&total)
	fmt.Println("")

	for i := 0; i < total; i++ {
		var name string
		fmt.Print("Input ", i+1, " Student’s Name :  ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input ", name, " Score :  ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is ", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : ", nameMax, " (", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : ", nameMin, " (", scoreMin, ")")
}
