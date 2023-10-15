# Summary

1. Entity Relationship Diagram (ERD) adalah diagram yang digunakan untuk perancangan suatu database dan menunjukan relasi antar objek atau entitas beserta atribut-atributnya secara detail.

<br>

2. Relasi database dibagi menjadi tiga yaitu :
   * one to one
   * one to many
   * many to many

<br>

3. Jenis perintah SQL dibagi tiga yaitu :
   * DDL (Data Definition Language)
   * DML (Data Manipulation Language)
   * DCL (Data Control Languange)

# Database, DDL, DML
## Database
Database adalah suatu sistem perangkat lunak yang dirancang untuk mengelola, menyimpan, dan mengatur kumpulan data yang terstruktur atau terorganisir. Dalam konteks pengembangan perangkat lunak dan manajemen informasi, database berperan penting dalam menyimpan dan mengakses data secara efisien. Ini adalah komponen kunci dari berbagai jenis aplikasi dan sistem, termasuk aplikasi bisnis, situs web, sistem manajemen konten, aplikasi perbankan, dan banyak lagi.

### Konsep dasar database
* **Data**: Database digunakan untuk menyimpan berbagai jenis data, seperti informasi pelanggan, produk, pesanan, atau catatan lainnya. Data ini dapat berupa teks, angka, gambar, suara, atau jenis data lainnya.

* **Struktur Data**: Data dalam database memiliki struktur yang terdefinisi dengan baik. Ini berarti data diatur dalam tabel, kolom, dan baris, yang memungkinkan untuk mengorganisasikan dan mengakses data dengan mudah.

* **Manajemen Data**: Database digunakan untuk mengelola data secara efisien. Ini termasuk operasi seperti penyimpanan, pengambilan, pembaruan, dan penghapusan data.

* **Integritas Data**: Database memiliki mekanisme untuk memastikan bahwa data tetap konsisten dan akurat. Ini mencakup penerapan aturan integritas referensial, unik, dan lainnya.

* **Pengindeksan Data**: Untuk meningkatkan kecepatan pencarian dan akses data, database menggunakan indeks untuk kolom tertentu.

## Data Definition Language (DDL)
Data Definition Language (DDL) dalah salah satu bagian dari SQL (Structured Query Language) yang digunakan untuk mendefinisikan struktur dan skema database. DDL digunakan untuk membuat, mengubah, dan menghapus tabel, indeks, kunci, dan objek database lainnya. Ini berfungsi sebagai bahasa untuk mendefinisikan dan mengelola struktur data dalam database.

### Perintah DDL yang umum digunakan
* **CREATE**: Untuk membuat database atau tabel
* **ALTER** : Mengubah struktur tabel yang sudah ada. Dapat menambahkan kolom baru, mengubah tipe data, atau menghapus kolom.
* **DROP** : Untuk menghapus database atau tabel

## Data Manipulation Language (DML)
Data Manipulation Language (DML) adalah salah satu bagian dari SQL (Structured Query Language) yang digunakan untuk mengelola dan memanipulasi data yang ada dalam tabel atau objek database. DML berfokus pada operasi seperti menyisipkan data baru, mengubah data yang sudah ada, menghapus data, dan mengambil data dari tabel. 

### Perintah DML yang umum digunakan 
* **INSERT**: Untuk menyisipkan data baru ke dalam tabel.
* **UPDATE**: Untuk mengubah data yang sudah ada dalam tabel.
* **DELETE**: Untuk menghapus data dari tabel.
* **SELECT**: Untuk mengambil data dari tabel.