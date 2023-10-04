# Summary

1. Clean architecture adalah arsitektur aplikasi dengan memisahkan setiap bagian kode sesuai dengan fungsinya yang bertujuan agar kode bersifat modular, mudah dikembangkan, mudah di perbaiki, dan mudah dalam melakukan pengujian.

<br>

2. Batasan yang diperlukan sebelum mendesain clean architecture yaitu :
   * Independent of framework, ini memungkinkan untuk menggunakan framework sebagai alat, daripada memiliki batasan dari framework tersebut.
   * Testable, business rules dapat melakukan pengujian tanpa melibatkan module yang lain
   * Independent of UI, tampilan dapat dengan mudah berubah, tanpa merubah fungsi dari sistem.
   * Independent database, business rules tidak bergantung penuh ke database sehingga dapat mengganti database dengan mudah
   * Independent of any external, business rules tidak perlu mengetahui sistem eksternal yang digunakan

<br>

3. Manfaat clean architecture yaitu
   * struktur standar, sehingga mudah dalam membangun sebuah aplikasi
   * Pengembangan yang lebih cepat dalam jangka panjang
   * Ketergantungan mocking menjadi hal yang sepele di unit test
   * pergantian yang mudah dari prototype ke solusi yang tepat(contoh pergantian in-memory storage ke SQL database)  

<br>

4. Layer-layer yang ada di clean architecture yaitu
   * Entities layer, berisi business object yang mencerminkan konsep dari dikelola aplikasi
   * Use case - domain layer, berisi business logic
   * Controller - presentation layer, berisi tampilan data dan menangani interaksi pengguna
   * Driver - data layer, mengelola data aplikasi, contoh mengambil data dari jaringan, mengelola data cache