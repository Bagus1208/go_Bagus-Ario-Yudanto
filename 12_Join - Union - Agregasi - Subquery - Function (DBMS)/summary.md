# Summary

1. Join berfungsi untuk menampilkan data dari beberapa tabel yang memiliki relasi sekaligus. Namun jika memakainya terlalu banyak dapat membuat performa semakin menurun. Join ada beberapa jenis yaitu :
   1. Inner join
   2. Left join
   3. Right join

<br>

2. Subquery adalah query yang berada didalam query lain. Fungsi subquery adalah untuk mengembalikan data yang digunakan untuk batasan data yang akan diambil dari query utama. Contoh subqury : 
    * SELECT * FROM products WHERE harga > (SELECT MIN(harga) FROM products);

<br>

3. Function bisa mengembalikan sebuah nilai pada pemanggilnya dan juga bisa menjalankan query dengan kondisi tertentu yang biasa disebut trigger function.