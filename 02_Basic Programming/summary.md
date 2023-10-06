# Summary

1. Golang adalah bahasa pemrograman yang dibuat oleh google yang rilis pertama kali ke publik pada tahun 2007 yang di kembangkan oleh Robert Grisemer, Rob Pike, Ken Thompson, Ian Lance Taylor, dan Rus Cox

<br>

2. Di Golang ada beberapa tipe data yaitu :
   1. Boolean yang bernilai True dan False
   2. Numeric yang bernilai angka
   3. String yang bernilai kumpulan karakter 

<br>

3. Ada beberapa control di Golang
   1. Untuk menangani percabangan menggunakan
      1. if
      2. switch
   2. untuk menangani perulangan menggunakan
      1. for

# Basic Programming

## Golang
Golang adalah bahasa pemrograman yang dibuat oleh google yang rilis pertama kali ke publik pada tahun 2007 yang di kembangkan oleh Robert Grisemer, Rob Pike, Ken Thompson, Ian Lance Taylor, dan Rus Cox

Contoh penggunaan golang :
1. Program aplikasi
   * E-commerce
   * Music Player
   * Social media app
2. Program sistem
   * APIs
   * Game engines
   * CLI apps

### Alasan menggunakan golang :
1. Simple
2. Gabungan dari beberapa bahasa pemrograman
3. Sintak yang ringan
4. Built-in concurrency
5. Open source
6. Digunakan oleh perusahaan besar

### Main syntak di golang
```
package main

import "fmt"

func main() {
   fmt.Println("Hello World")
}
```
## Variable dan Tipe Data
Variabel adalah tempat untuk menyimpan sebuah informasi yang dideskripsikan dengan sebuah nama dan memiliki tipe data

Tipe data di golang :
1. Boolean
   * true
   * false
2. Numeric
   * integer
     * uint8   (0 to 255)
     * uint16  (0 to 65535)
     * uint32  (0 to 4294967295)
     * uint64  (0 to 18446744073709551615)
     * uint    (same as uint32 or uint64)
     * int8    (-128 to 127)
     * int16   (-32768 to 32767)
     * int32   (-2147483648 to 2147483647)
     * int64   (-9223372036854775808 to 9223372036854775807)
     * int     (same as int32 or int64)
     * rune    (same as int32)
     * byte    (same as int8)
   * float
     * float32 
     * float64
   * complex
     * complex64
     * complex128

### Deklarasi variabel di golang
1. Cara panjang
```
var angka int
var angka int = 8
var firstname, lastname string
var firstname, lastname string = "Bagus", "Yudanto"
```
2. Cara pendek
```
name := "Bagus"
firstname, lastname := "Bagus", "Yudanto"
```

### Zero value di golang :
1. boolean = false
2. float = 0.0
3. integer = 0
4. string = ""

## Konstanta
Konstanta adalah variabel yang isi nilainya tidak dapat diubah setelah dideklarasikan. Saat deklarasi konstanta wajib langsung memberikan nilainya.

### Deklarasi konstanta di golang
1. Satu konstanta
```
const Pi float64 = 3.14
```
2. Lebih dari satu konstanta
```
const (
   Pi float64 = 3.14
   Age int = 10
)
```

## Operators
### Aritmatika
1. Pertambahan (+)
2. Pengurangan (-)
3. Perkalian (*)
4. Pembagian (/)
5. Modulo (%)
6. Increment (++)
7. Decrement (--)

### Perbandingan
1. Sama dengan (==)
2. Tidak sama dengan (!=)
3. Lebih dari (>)
4. Kurang dari (<)
5. Lebih dari sama dengan (>=)
6. Kurang dari sama dengan (<=)

### Logika
1. AND (&&)
2. OR (||)
3. NOT (!)

### Bitwise
1. (&)
2. (|)
3. (^)
4. (<<)
5. (>>)

### Assigment
1. (=)
2. (+=)
3. (-=)
4. (*=)
5. (/=)
6. (%=)
7. (<<=)
8. (>>=)
9. (&=)
10. (^=)
11. (|=)

### Miscellaneous (untuk pointer)
1. (&)
2. (*)

## Struktur kontrol
### IF, ELSE IF, ELSE
```
hour := 15
if hour < 12 {
   fmt.Println("Selamat Pagi")
} else if hour < 18 {
   fmt.Println("Selamat Sore")
} else {
   fmt.Println("Selamat Malam")
}
```

### Switch
```
today := 2
switch hour {
case 1: 
   fmt.Println("Hari ini hari Senin")
case 2:
   fmt.Println("Hari ini hari Selasa)
default:
   fmt.Println("Salah input")
}
```

### Switch without expression
```
today := 2
switch {
case today == 1: 
   fmt.Println("Hari ini hari Senin")
case today == 2:
   fmt.Println("Hari ini hari Selasa)
default:
   fmt.Println("Salah input")
}
```

### Fallthrough
Code :
```
value := 42
switch value {
case 100:
   fmt.Println(100)
   fallthrough
case 42:
   fmt.Println(42)
   fallthrough
case 1: 
   fmt.Println(1)
   fallthrough
default:
   fmt.Println("default")
}
```

Output :
```
42
1
default
```

### Loop
```
for i := 0; i < 10; i++ {
   fmt.Println(i)
}
```