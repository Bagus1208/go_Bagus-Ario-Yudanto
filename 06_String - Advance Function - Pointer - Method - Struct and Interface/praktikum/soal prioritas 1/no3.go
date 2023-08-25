package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var result bool

	if len(a) < len(b) {
		a, b = b, a
	}

	result = strings.Contains(a, b)

	if result {
		return b
	}
	return "Tidak ada substring yang sama"

}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA
	fmt.Println(Compare("Yudanto", "Yuda"))   //Yuda
}
