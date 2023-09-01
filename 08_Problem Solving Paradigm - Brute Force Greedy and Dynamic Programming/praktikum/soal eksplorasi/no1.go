package main

import "fmt"

func ToRoman(number int) {
	var ones = map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX"}
	var tens = map[int]string{1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC"}
	var hundreds = map[int]string{1: "C", 2: "CC", 3: "CCC", 4: "CD", 5: "D", 6: "DC", 7: "DCC", 8: "DCCC", 9: "CM"}
	var thousands = map[int]string{1: "M", 2: "MM", 3: "MMM", 4: "MMMM"}

	var ribuan = thousands[number/1000]
	var ratusan = hundreds[(number%1000)/100]
	var puluhan = tens[(number%100)/10]
	var satuan = ones[(number % 10)]

	fmt.Println(ribuan + ratusan + puluhan + satuan)
}

func main() {
	ToRoman(4)
	ToRoman(9)
	ToRoman(23)
	ToRoman(2021)
	ToRoman(1646)
	ToRoman(2002)
}
