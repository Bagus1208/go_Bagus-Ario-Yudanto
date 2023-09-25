# Summary

1. Middleware adalah sebuah kode yang dijalankan sebelum/setelah proses request/respone. Cara kerja middleware seperti penjaga di pintu masuk mall yang mengecek setiap pengunjung sebelum masuk ke dalam mall.

<br>

2. Ada dua tipe setup di middleware echo yaitu :
   * echo#pre()    : middleware akan dijalankan sebelum router memproses request
   * echo#user()   : middleware akan dijalankan setelah router memproses request dan mempunyai akses penuh ke API

<br>

3. Beberapa middleware yang sering digunakan yaitu :
   * CORS middleware, untuk memungkinkan API diakses dengan alamat/domain yang berbeda
   * Rate limiter middleware, untuk membatasi jumlah request ke server
   * Logging middleware, untuk mencatat informasi http request 
   * Authentication middleware (basic auth dan JWT), untuk mengidentifikasi user yang melalukan request
