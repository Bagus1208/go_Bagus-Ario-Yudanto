package main

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (mobil *Mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

func (mobil *Mobil) tambahKecepatan(penambahan int) {
	mobil.KecepatanPerJam += penambahan
}

func main() {
	var mobilCepat = Mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	var mobilLamban = Mobil{}
	mobilLamban.berjalan()
}
