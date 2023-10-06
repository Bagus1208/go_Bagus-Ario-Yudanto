# Summary

1. Package strings berfungsi untuk memanipulasi tipe data string untuk berbagai keperluan. Contoh untuk membuat seluruh karakter pada variabel string menjadi huruf besar.

<br>

2. Ada beberapa jenis fungsi (function) yaitu :
   * Variadic function
      - variadic function digunakaan pada saat jumlah input parameneter belum diketahui.
   * Anonymous function
      - anonymous function adalah fungsi yang tidak memiliki nama. Berguna saat ingin membuat fungsi inline.
   * Closure
      - Closure adalah tipe khusus dari anonymous function yang mendeklarasikan variabel referensi di luar fungsi itu sendiri.
   * Defer function
      - Defer function adalah fungsi yang hanya dijalankan setelah parent function nya selesai dijalankan.

<br>

3. Pointer adalah variabel yang menyimpan alamat memori variabel lain. Dengan begitu jika kita merubah nilai suatu variabel, maka nilai variabel pointer yang mengacu ke alamat memori variabel tersebut akan berubah juga.

<br>

4. Struct adalah tipe yang ditentukan pengguna yang berisi kumpulan field/properti atau method. Method adalah fungsi yang melekat pada struct.

<br>

5. Interface adalah kumpulan signatured method yang dapat diimplementasikan oleh suatu objek. Karenanya interface mendefinisikan perilaku objek.


# String - Advance - Function - Pointer - Method - Struct and Interface
## String
String adalah tipe data yang menyimpan karakter atau kalimat.
### Mengetahui panjang string
```
var name = "Bagus"
fmt.Println(len(name)) // output : 5
```

### Membandingkan dua string
```
var str1 = "abc"
var str2 = "cba"
fmt.Println(str1 == str2) // output : false
```

### Mengecek kandungan kata
```
var str = "something"
var substr = "some"
res := strings.Contains(str, substr)
fmt.Println(res) // output true
```

### Mengambil huruf dalam sebuah kata/kalimat
```
var value = "cat;dog"
substring := value[4:len(value)]
fmt.Println(substring) // output : dog
```

### Mengganti karakter yang ada di string
```
s := "this[thing]I would like to remove"
t := strings.Replace(s, "[things]", " ", -1)
fmt.Printf("%s\n", t) // output : this I would like to remove
```

### Menyisipkan kata 
```
p := "Bagus"
index := 2
q := p[:index] + "ng a" + p[index:]
fmt.Println(q) // output : Bang agus
```

## Function
Function dalam pemrograman adalah blok kode yang memiliki nama dan dapat digunakan untuk menjalankan tugas-tugas tertentu. Fungsi digunakan untuk mengorganisasi kode ke dalam unit yang lebih kecil dan terkelola, yang dapat dipanggil (dipakai) berulang-ulang dalam program.

### Variadic function
Beberapa kondisi penggunaan variadic function :
* Untuk menggantikan slice sebagai parameter yang dimasukkan ke function
* Ketika jumalh input tidak diketahui
* Untuk meningkatkan keterbacaan

Contoh variadic function
```
// membuat function sum dengan variadic function
func sum(numbers ...int) int {
   var total = 0 
   
   for _, number := range numbers {
      total += number
   }

   return total
}

// cara menggunakan variadic function
avg := sum(3, 4, 6, 2) // dapat memasukkan argumen berapapun
fmt.Println(avg)
```

### Anonymous function / literal function
Anonymous function adalah function yang tidak memiliki nama. Ini berguna ketika kita ingin membuat function dalam satu baris.

Contoh anonymous function
```
func() {
   fmt.Println("This is anonymous function")
}()

// memasukkan function kedalam variabel
value := func() {
   fmt.Println("This is function inside variable)
}

// cara memanggil function dalam variabel
value()

// anonymous function dengan parameter
func(name string) {
   fmt.Println("Hello", name)
}("Bagus")
```

## Closure
Closure adalah tipe spesial dari anonymous function yang mereferensikan variabel yang dideklarasikan di luar fungsi itu sendiri

Contoh closure 
```
func newCounter() func() int {
   count := 0
   return func() int {
      count += 1
      return count
   }
}

func main()
counter := newCounter()
fmt.Println(counter()) // output : 1
fmt.Println(counter()) // output : 2
```

## Defer 
Defer adalah function yang dijalankan paling terakhir didalam sebuah kode. Jika ada lebih dari satu defer maka akan berjalan seperti stack, satu per satu

Contoh defer
```
func main() {
   defer func() {
      fmt.Println("Bagus")
   }()

   fmt.Println("Hello")
}
```

## Pointer
Pointer adalah variabel yang menyimpan alamat memori variabel lain. Pointer dapat mengubah nilai pada variabel yang ditunjuk.

### Dua operator penting di pointer
1. operator *: untuk mendeklarasi variabel pointer dan mengakses nilai yang disimpan didalam memori
2. operator &: mengembalikan alamat memori variabel dan akses alamat variabel untuk pointer

### Deklarasi pointer
```
var name string = "Bagus"
var nameAddress *string = &name

fmt.Println(name)          // Bagus
fmt.Println(&name)         // 0x000010050
fmt.Println(nameAddress)   // 0x000010050
fmt.Println(*nameAddress)  // Bagus

var pointer = new(string)
pointer = nameAddress
fmt.Println(pointer)       // 0x000010050
fmt.Println(*pointer)      // Bagus
```

### Mengubah nilai variabel yang ditunjuk
```
var name string = "Bagus"
var nameAddress *string = &name

*nameAddress = "Ario"

fmt.Println(name)    // Ario
```

## Struct
Struct adalah tipe data kompleks yang digunakan untuk mengelompokkan beberapa variabel dengan tipe data yang berbeda menjadi satu kesatuan logis. Struct dapat digunakan untuk merepresentasikan entitas atau objek yang memiliki sejumlah atribut atau properti.

### Deklarasi struct
```
type person struct {
   Name     string
   Age      int
   Address  strint
}

var bagus = person{
   Name: "Bagus",
   Age: "21",
   Address: "Cibinong",
}
```

## Method
Method adalah function yang melekat pada suatu tipe (biasanya struct atau tipe data lain)

### Deklarasi method
```
func (receiver structType) MethodName(parameter) return {

} 
```

### Alasan menggunakan method daripada function
* Membantu kita menulis gaya berorientasi objek di golang
* Method membantu kita menghindari konflik penamaan
* Pemanggilan method lebih mudah untuk dibaca dan dimengerti daripada pemanggilan function

## Interface
Interface adalah kumpulan method signatur yang dapat diimplementasikan oleh suatu objek. 

### Deklarasi interface
```
type interface_name interface {
   method_name1 <return_type>
   method_name2 <return_type>
   method_name3 <return_type>
   ...
}
```

Contoh struct yang mengimplementasi interface
```
type calculate interface {
   large() int
}

type square struct {
   side int
}

func (s square) large() int {
   return s.side * s.side
}

func main() {
   var dimResult calculate
   dimResult = square{10}
   fmt.Println(dimResult.large())   // output : 100
}
```

### Interface kosong untuk nilai yang dinamis
```
func describe(i interface{}) {
   fmt.Println("(%v, %T)\n", i, i)
}

func main() {
   var i interface{}
   describe(i)          // (<nil>, <nil>)

   i = 42
   describe(i)          // (42, int)

   i = "hello"
   describe(i)          // (hello, string)
}
```

### Type assertion
```
var secret interface{}

secret = 2
var number = secret.(int) * 10
fmt.Println(number)  // 20

secret = []string{"apple", "mango", "strawberry"}
var fruits = strings.Join(secret.([]string), ", ")
fmt.Println(fruits)  // apple, mango, strawberry
```

## Package
Package dalah cara untuk mengorganisir dan mengelompokkan kode program menjadi unit-unit terpisah yang dapat digunakan kembali. Setiap berkas sumber golang biasanya dimulai dengan pernyataan package, yang menentukan nama paket yang digunakan untuk mengelompokkan kode di dalamnya. Paket adalah salah satu konsep paling penting dalam golang dan berfungsi sebagai fondasi struktur proyek golang.

Contoh package
```
package aritmatika

func Tambah(a, b int) int {
   return a + b
}

func Kurang(a, b int) int {
   return a - b
}
```
```
package main

import (
   "aritmatika"
   "fmt"
)

func main() {
   fmt.Println(aritmatika.Tambah(5, 10))  // 15
}
```

Cara membuah variabel dan function dapat diakses dari luar package adalah dengan menggunakan huruf kapital pada huruf awal di variabel atau function

contoh 
```
var Name string   // dapat diakses dari luar package
var name string   // tidak dapat diakses dari luar package

func Tambah() int // dapat diakses dari luar package
func tambah() int // tidak dapat diakses dari luar package
```

## Error Handling
Error handling adalah bagian penting dari pemrograman di golang. Golang memiliki pendekatan yang cukup unik dalam penanganan kesalahan, yang didasarkan pada penggunaan nilai yang disebut "error" sebagai nilai yang dapat dikembalikan dari fungsi-fungsi yang mungkin gagal. 

Berikut adalah cara penanganan kesalahan di golang :
* **Fungsi Mengembalikan Nilai Error**: Dalam golang, banyak fungsi mengembalikan dua nilai, yaitu hasil operasi dan error (jika ada). Error adalah tipe data yang memungkinkan Anda untuk mengecek apakah suatu operasi berhasil atau gagal. Biasanya, jika operasi berhasil, nilai error akan nil atau nil, dan jika operasi gagal, nilai error akan berisi informasi tentang kesalahan tersebut.

   Contoh penggunaan error dalam sebuah fungsi:
   ```
   func divide(a, b float64) (float64, error) {
      if b == 0 {
        return 0, errors.New("Pembagian oleh nol tidak diizinkan")
      }
      return a / b, nil
   }
   ```

* **Pemeriksaan Error**: Setelah pemanggilan fungsi yang mungkin mengembalikan error, Anda perlu memeriksa error tersebut. Ini dilakukan dengan menggunakan pernyataan kondisional seperti if untuk memeriksa apakah error adalah nil atau tidak. Jika error bukan nil, itu berarti ada kesalahan yang perlu ditangani.

   Contoh pemeriksaan error:
   ```
   result, err := divide(10, 0)
   if err != nil {
      fmt.Println("Error:", err)
   } else {
      fmt.Println("Hasil:", result)
   }
   ```

* **Penggunaan panic dan recover**: Di golang, Anda juga memiliki opsi untuk menghentikan eksekusi program dengan panic dan menangkapnya dengan recover. Ini digunakan untuk situasi yang sangat fatal, seperti ketika terjadi kesalahan yang tidak dapat diperbaiki dan program harus segera berhenti.

   Contoh panic dan recover:
   ```
   defer func() {
        if r := recover(); r != nil {
            fmt.Println("Program panic:", r)
        }
    }()
    // Penyebab panic
    panic("Ini adalah kesalahan serius")
   ```

* Paket errors: Golang memiliki paket bawaan errors yang digunakan untuk membuat dan mengelola error. Anda dapat menggunakan errors.New() untuk membuat error baru atau mengembalikan error yang sudah ada dalam paket ini.

   Contoh penggunaan paket errors:
   ```
   import "errors"

   func divide(a, b float64) (float64, error) {
      if b == 0 {
         return 0, errors.New("Pembagian oleh nol tidak diizinkan")
      }
      return a / b, nil
   }
   ```
