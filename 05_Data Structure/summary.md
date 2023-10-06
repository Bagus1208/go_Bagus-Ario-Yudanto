# Summary

1. Ada 3 struktur data di golang yaitu : 
   1. Array
      - Array dapat menampung data yang bertipe sama dengan jumlah yang sudah ditentukan. Data dalam array dapat diakses menggunakan indek. 
   2. Slice
      - Slice sama seperti array, tetapi jumlah data yang dapat disimpan dapat berubah atau dinamik
   3. Map
      - Kumpulan pasangan data yang berisi kunci (key) dan nilai (value). Kunci di map bersifat unik sehingga didalam map tidak akan ada kunci yang sama. 

<br>

2. Fungsi (function) adalah potongan kode dalam sebuah blok yang dipanggil untuk melakukan tugas tertentu. Fungsi dapat membuat kode menjadi lebih ringkas, mudah dibaca, dan rapih.  

<br>

3. Dengan adanya fungsi, kita tidak perlu menulis kode yang sama berulang kali untuk mendapatkan jawaban yang sama karena fungsi dapat dipanggil berkali-kali sesuai dengan keinginan

# Struktur Data
## Array
Array adalah struktur data yang mengandung kumpulan data yang bertipe sama dengan ukuran alokasi tetap.

### Deklarasi array
1. Tanpa ada isinya
   ```
   var books [2]string
   ```
2. Dengan isinya
   ```
   var books [2]string = [2]string{"Edensor", "Laskar Pelangi", "Sapiens"}
   ```

### Memasukkan dan mengakses nilai array
1. Memasukkan nilai array
   ```
   books[0] = "Edensor"
   ```
2. Mengakses nilai array
   ```
   fmt.Println(books)      // mengakses seluruh nilai dalam array
   fmt.Println(books[0])   // mengakses nilai pertama dalam array
   ```
3. Mengakses nilai array dengan looping
   ```
   // cara 1
   for i := 0; i < len(books); i++ {
      fmt.Println(books[i])
   }

   // cara 2
   for index, book := range books {
      fmt.Println(index, " = ", book)
   }

   for _, book := range books {
      fmt.Println(book)
   }

   // cara 3
   index := 0
   for range books {
      fmt.Println(books[index])
      index++
   }
   ```

## Slice
Slice merupakan struktur data yang mengandung kumpulan data yang bertipe sama, slice sama seperti array namun yang membedakan adalah ukuran alokasinya dinamis.

### Deklarasi slice 
1. Membuat slice dari array
   ```
   var books [2]string = [2]string{"Edensor", "Laskar Pelangi", "Sapiens"}

   var slice_books []string = books[0:2]
   fmt.Println(slice_books)

   // output ["edensor", "Laskar Pelangi"]
   var books [2]string = [2]string
   ```
2. Tanpa ada isinya
   ```
   var books []string
   ```
3. Dengan isinya
   ```
   var books []string = []string{"Edensor", "Laskar Pelangi", "Sapiens"}
   ```
4. Menggunakan keyword make()
   ```
   var books = make([]string, 2, 4)
   ```
   note :
   * parameter pertama menandakan tipe slice
   * parameter kedua menandakan panjang slice
   * parameter ketiga menandakan kapasitas slice

### Menambahkan nilai slice
```
var books []string = []string{"Edensor", "Laskar Pelangi", "Sapiens"}

books = append(books, "Atomic Habits")
```
Jika data yang disimpan kedalam slice sudah mencapai kapasitas maksimal lalu menambahkan nilainya lagi ke slice tersebut, maka kapasitas slice akan bertambah

### Copy slice
```
var books []string = []string{"Edensor", "Laskar Pelangi", "Sapiens"}

var copy_books = make([]string, 3, 6)

copy(copy_books, books)
```
note :
jika panjang slice yang akan mengcopy lebih kecil dari panjang slice yang akan dicopy, maka nilai yang dicopy sesuai dengan panjang slice yang mengcopy.

## Map
Map adalah struktur data yang menyimpan data berupa pasangan key dan value dimana key bersifat unik.

### Deklarasi map
1. Tanpa ada isinya
   ```
   var salary map[string]int
   ```
2. Dengan isinya
   ```
   var salary = map[string]int{
      "Umam": 1000,
      "Iswanul": 2000,
   }
   ```
3. Dengan keyword make()
   ```
   var salary = make(map[string]int)
   ```

### Memasukkan dan mengakses nilai map
1. Memasukkan nilai map
   ```
   var salary map[string]int

   salary["Bagus"] = 5000
   ```
2. Mengakses nilai map
   ```
   fmt.Println(salary["Bagus"])
   ```
   