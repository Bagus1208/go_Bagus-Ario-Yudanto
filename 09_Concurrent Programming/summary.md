# Summary 

1. Concurrent programming adalah program yang menjalankan beberapa tugas bergantian secara independen. Implementasi concurrent programming di golang adalah goroutine.

<br>

2. Perbedaan antara sequential, parallel, dan concurrent
   * Sequential
      - sebelum tugas baru dijalankan, tugas yang sedang dijalankan harus selesai dahulu.
   * Parallel
      - beberapa tugas dapat dijalankan bersamaan.
   * Concurrent
      - Beberapa tugas dijalankan secara bergantian hingga selesai.

<br>

3. Pada concurrent programming terkadang muncul masalah yaitu race condition dimana 2 thread atau lebih mengakses memory yang sama dalam waktu yang sama. Untuk mengatasi itu ada beberapa cara yaitu dengan menggunakan :
   * waitgroup
   * channel blocking
   * mutex

# Concurrent Programming
## Big search website
Big search website adalah website yang memungkinkan penggunanya untuk mencari informasi di internet. Biasanya memiliki indeks situs web dan konten web lainnya yang besar, dan menggunakan berbagai algoritma untuk menentukan peringkat hasil query penelusuran.

## Sequential program 
Sequential program adalah jenis program komputer yang dieksekusi secara berurutan, langkah demi langkah, dari awal hingga akhir tanpa cabang atau percabangan eksekusi.

## Parallel program
Parallel program adalah jenis program komputer yang dirancang untuk dieksekusi secara paralel, yang berarti bahwa beberapa instruksi atau tugas dapat dijalankan secara bersamaan pada beberapa unit pemrosesan atau inti CPU yang tersedia. Dalam program paralel, tugas-tugas yang berbeda dapat dikerjakan secara bersamaan untuk mengoptimalkan penggunaan sumber daya komputasi dan mempercepat eksekusi program.

## Concurrent program
Concurrent program adalah program yang dapat dijalankan secara bersamaan dengan program lain pada waktu yang sama. Program ini biasanya digunakan untuk tugas-tugas yang membutuhkan waktu lama dan memakan banyak sumber daya.
Di concurrent program, beberapa pekerjaan dapat dieksekusi secara independen dan mungkin muncul bersamaan.

## Goroutine
Goroutine adalah function atau method yang berjalan concurrently(independen) bersamaan dengan function atau method lain. Biaya pembuatan goroutine sangat kecil jika dibandingkan dengan thread. thread adalah proses yang ringan, atau dengan kata lain thread adalah unit yang mengeksekusi kode di bawah program. 

Contoh program yang menggunakan goroutine :
```
package main

import (
	"fmt"
	"sync"
	"time"
)

func MultipleNumber(number int, wait *sync.WaitGroup) {
	defer wait.Done()

	for i := number; i <= number*10; i += number {
		time.Sleep(3 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	var wait sync.WaitGroup

	wait.Add(1)
	go MultipleNumber(3, &wait)

	wait.Wait()
	fmt.Println("DONE")
}
```

## Channel 
Channel adalah struktur data yang digunakan untuk mengirim dan menerima data antara goroutine secara aman dan sinkron. Channel memungkinkan goroutine berkomunikasi dan berkoordinasi satu sama lain dalam aplikasi yang paralel atau concurrent.

### Jenis-jenis channel
1. **Unbufferd channel**: ketika mengirim data ke unbuffered channel (dengan menggunakan operator <-), pengirim akan terblokir sampai ada penerima yang siap menerima data tersebut. Ini berarti pengiriman dan penerimaan data melalui unbuffered channel terjadi secara bersamaan (synchronous). 

	Contoh pembuatan unbuffered channel : 
	```
	c := make(chan string)
	```
2. **Buffered channel**: Ketika Anda mengirim data ke buffered channel, pengirim tidak akan terblokir jika kapasitas channel belum tercapai. Ini berarti pengiriman dapat dilakukan secara asinkron (tidak selalu terblokir).
Pengiriman data ke buffered channel akan terblokir hanya jika kapasitas sudah tercapai, dan penerima belum membaca data dari channel.

	Contoh pembuatan buffered channel :
	```
	c := make(chan int, 3) // parameter kedua menentukan kapasitas penyimpanan channel
	```


### Contoh program menggunakan channel
```
package main

import (
	"fmt"
	"sync"
)

func MultipleOf3(channel chan<- int, wait *sync.WaitGroup) {
	defer wait.Done()

	for i := 3; i <= 30; i += 3 {
		channel <- i
	}
	close(channel)
}

func PrintChannel(channel <-chan int, wait *sync.WaitGroup) {
	defer wait.Done()

	for number := range channel {
		fmt.Println(number)
	}
}

func main() {
	var channel = make(chan int)
	var wait sync.WaitGroup

	wait.Add(2)
	go MultipleOf3(channel, &wait)
	go PrintChannel(channel, &wait)

	wait.Wait()
	fmt.Println("\nDONE")
}
```

## Select 
Select dapat mempermudah untuk mengotrol data komunikasi dari satu atau lebih channel

contoh penggunaan select :
```
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Pesan dari goroutine 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Pesan dari goroutine 2"
	}()

	// Menggunakan select untuk menerima data dari channel yang siap
	select {
	case msg1 := <-ch1:
		fmt.Println("Menerima:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Menerima:", msg2)
	}

	fmt.Println("Selesai")
}
```

## Race condition
Race condition adalah situasi dalam pemrograman konkurensi atau paralel di mana perilaku program menjadi tidak dapat diprediksi karena dua atau lebih goroutine atau thread bersaing untuk mengakses dan memodifikasi sumber daya bersama (seperti variabel atau struktur data) tanpa koordinasi yang memadai. Dalam race condition, hasil akhir dari eksekusi program tergantung pada urutan eksekusi dan akses bersama yang tidak terduga, yang dapat menghasilkan hasil yang tidak konsisten atau bahkan tidak benar.

Contoh race condition :
```
var x = 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x += 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter :", x)
```

## Fixing data race
1. Bloking dengan WaitGroups
	```
	func RunAsynchronous(group *sync.WaitGroup) {
		defer group.Done()

		group.Add(1)

		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}

	func main() {
		group := &sync.WaitGroup{}

		for i := 0; i < 100; i++ {
			go RunAsynchronous(group)
		}

		group.Wait()
		fmt.Println("Done")
	}
	```
2. Bloking dengan channel
	```
	func main() {
		var data int
		var wg sync.WaitGroup
		ch := make(chan struct{}, 1) // Buat blocking channel dengan kapasitas 1

		wg.Add(2)

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				ch <- struct{}{} // Mengirim sinyal ke channel (mengambil tempat di channel)
				data++
				<-ch // Mengambil sinyal dari channel (memberi tempat di channel)
			}
		}()

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				ch <- struct{}{} // Mengirim sinyal ke channel (mengambil tempat di channel)
				data++
				<-ch // Mengambil sinyal dari channel (memberi tempat di channel)
			}
		}()

		wg.Wait()
		fmt.Println("Nilai data akhir:", data)
	}
	```
4. Menggunakan mutex
	```
	func main() {
		var x = 0
		var mutex sync.Mutex

		for i := 1; i <= 1000; i++ {
			go func() {
				for j := 1; j <= 100; j++ {
					mutex.Lock()
					x += 1
					mutex.Unlock()
				}
			}()
		}
		time.Sleep(5 * time.Second)
		fmt.Println("Counter :", x)
	}
	```