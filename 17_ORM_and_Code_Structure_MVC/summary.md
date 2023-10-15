# Summary

1. ORM (Object-Relational Mapping) adalah teknik pemrograman untuk mengubah data dari bentuk tabel di relational database ke bentuk object. Contoh ORM di golang adalah GORM.

<br>

2. Keuntungan dan kerugian ORM yaitu
   * Keuntungan :
     * Mengurangi query yang berulang
     * Otomatis mengambil data ke objek yang siap pakai
     * Cara mudah jika ingin menyaring data sebelum menyimpan ke database
     * Ada beberapa ORM yang memiliki fitur query cache
  
   * kekurangan
     * Perlu mempelajari kebiasaan ORM terlebih dahulu sebelum menggunakannya
     * Mengurangi performa proses karena ORM memuat data yang tidak diperlukan
     * Tidak bisa menggunakan fungsi query yang terlalu spesifik

<br>

3. Database migration adalah suatu cara untuk memperbarui versi database agar sesuai dengan perubahan versi aplikasi. Perubahan dapat berupa peningkatan ke versi terbaru atau kembali ke versi sebelumnya.

<br>

4. Database relasional di GORM
   * Belongs to
   * Has one
   * Has many
   * Many to many

<br>

5. Keuntungan database migration
   * Mudah untuk update database
   * Mudah untuk rollback database
   * Dapat tracking struktur database
   * Riwayat struktur database ditulis pada sebuah kode
   * Selalu kompatibel dengan perubahan versi aplikasi

<br>

5. MVC adalah singkatan untuk Model, View, Controller. Tujuan MVC adalah untuk membagi kode berdasarkan fungsinya sehingga memudahkan dalam pengembangan kode.