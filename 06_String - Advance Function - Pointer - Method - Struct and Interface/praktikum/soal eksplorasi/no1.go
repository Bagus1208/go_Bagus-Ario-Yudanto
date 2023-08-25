package main

import (
	"fmt"
	"unicode"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var (
		nameEncode string
		temp       byte
	)

	for i := 0; i < len(s.name); i++ {
		if s.name[i] == 32 {
			nameEncode += string(s.name[i])
			continue
		}
		if unicode.IsUpper(rune(s.name[i])) {
			temp = s.name[i] - 65
			nameEncode += string(90 - temp)
		} else {
			temp = s.name[i] - 97
			nameEncode += string(122 - temp)
		}
	}

	s.nameEncode = nameEncode

	return nameEncode
}

func (s *student) Decode() string {
	var (
		nameDecode string
		temp       byte
	)
	for i := 0; i < len(s.nameEncode); i++ {
		if s.nameEncode[i] == 32 {
			nameDecode += string(s.nameEncode[i])
		}
		if unicode.IsUpper(rune(s.nameEncode[i])) {
			temp = s.nameEncode[i] - 90
			nameDecode += string(65 - temp)
		} else {
			temp = s.nameEncode[i] - 122
			nameDecode += string(97 - temp)
		}
	}

	return nameDecode
}

func main() {

	// var nama = student{}
	// var key Chiper = &nama
	// var nama string
	// var alamat string

	// fmt.Print("Masukkan nama Anda : ")
	// scanner.Scan()
	// nama = scanner.Text()
	// fmt.Print("Masukkan alamat : ")
	// scanner.Scan()
	// alamat = scanner.Text()
	// fmt.Println("Hallo", nama, alamat)

	// fmt.Print("Input Nama : ")
	// fmt.Scan(&nama.name)
	// fmt.Println(nama.Encode())
	// fmt.Println(nama.nameEncode)
	// fmt.Println(nama.Decode())

	var menu int
	var a student = student{}
	var c Chiper = &a

	for {
		var choice string

		fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
		fmt.Scan(&menu)

		if menu == 1 {
			fmt.Print("\nInput Student Name: ")
			fmt.Scan(&a.name)
			fmt.Println("\nEncode of student’s name " + a.name + " is : " + c.Encode())
		} else if menu == 2 {
			fmt.Print("\nInput Student Name: ")
			fmt.Scan(&a.name)
			fmt.Println("\nDecode of student’s encode name " + a.nameEncode + " is : " + c.Decode())
		}

		fmt.Print("Run the program again? (yes/no) : ")
		fmt.Scan(&choice)

		if choice == "no" {
			break
		}
		fmt.Println("\n")
	}
}
