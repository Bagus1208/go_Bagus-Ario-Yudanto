# Unit Testing

## Software Testing
Software testing adalah suatu proses menganalisis item perangkat lunak untuk mendeteksi antara kondisi yang ada dan yang diperlukan dan untuk mengevaluasi fitur item perangkat lunak.

### Tujuan dari testing yaitu :
   * Mencegah fungsi yang semula berjalan secara tiba-tiba berhenti bekerja
   * Lebih yakin saat melakukan refactoring 
   * Memperbaiki desain code
   * Memudahkan dalam membuat dokumentasi kode
   * Memudahkan dalam bekerja tim

### Tingkatan dari testing yaitu :
   * UI testing/end-to-end testing, yaitu menguji secara keseluruhan aplikasi
   * Integration testing, yaitu menguji module secara spesifik atau sub-system melalui API
   * Unit testing, yaitu menguji bagian paling kecil sebuah kode (biasanya fungsi atau method) 

## Framework
Terdapat framework yang menyediakan alat untuk melakukan testing secara efisien, contohnya di golang ada testify

## Struktur Unit Testing
Ada dua pola pembuatan testing yang biasa digunakan yaitu :
   * Menyatukan semua file test didalam folder tests
   * Menyimpan file test satu folder dengan file production

## Runner
* Alat yang digunakan untuk menjalankan file testing
* Gunakan mode tontonan (jika ada perubahan di kode, secara otomatis akan menjalankan testing, membuat TDD lebih efisien)
* Pilih runner yang dapat berjalan paling cepat

## Mocking
Mocking adalah objek tiruan yang menirukan perilaku dari objek yang asli.

## Coverage
Programmer perlu memastikan bahwa mereka telah membuat testing yang cukup. Coverage tool dapat menganalisa aplikasi kode saat menjalankan testing.

### Coverage report
* Function
* Statement
* Branch
* Line

### Report format
* CLI
* XML
* HTML
* LCOV

## Melakukan Testing di Golang
### Membuat file test baru
1. Nama file diakhiri dengan _test. Contoh user_test.go
2. Lokasi file :
   * Folder yang sama, package yang sama
   * Folder yang sama, package berbeda
  
### Menulis fungsi test
1. Nama : Test dan huruf kapital
2. Harus mempunyai fungsi signatur func TestXxx(t *testing.T)

### Menjalankan file test
```
go test ./lib/calculate -cover
```

### Menjalankan file test dengan report coverage
```
go test ./lib/calculate -coverprofile=cover.out && go tool cover -html=cover.out
```