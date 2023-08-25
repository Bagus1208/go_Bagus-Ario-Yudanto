package main

import "fmt"

type Car struct {
	CarType string
	Fuelin  float64
}

func (car Car) EstimatedMileage() float64 {
	var mile float64

	mile = car.Fuelin * 1.5

	return mile
}

func main() {
	var myCar Car

	fmt.Print("Masukkan tipe mobil Anda : ")
	fmt.Scanln(&myCar.CarType)
	fmt.Print("Berapa liter bahan bakar mobil Anda : ")
	fmt.Scanln(&myCar.Fuelin)

	fmt.Println("\nTipe mobil Anda :", myCar.CarType)
	fmt.Println("Estimasi jarak tempuh :", myCar.EstimatedMileage(), "mile")
}
