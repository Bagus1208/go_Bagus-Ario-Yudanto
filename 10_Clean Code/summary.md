# Summary

1. Clean code adalah istilah untuk penulisan kode program yang mudah dibaca, dipahami, dan diubah oleh pembuat ataupun orang lain.

<br>

2. Clean code sangat penting untuk :   
   * work collaboration
   * feature development
   * faster development

<br>

3. Ada beberapa prinsip clean code, yaitu :
   * KISS (Keep It So Simple)
   * DRY (Don't Repeat Yourself)
   * Refactoring

# Clean Code
Clean code adalah konsep dalam pengembangan perangkat lunak yang berfokus pada menulis kode yang mudah dibaca, dimengerti, dan dapat dipelihara oleh pengembang lain atau diri sendiri di masa depan. Prinsip clean code didasarkan pada pemikiran bahwa kode sumber adalah komunikasi antara pengembang, dan kode yang baik seharusnya seperti buku yang bagus, mudah dimengerti oleh pembaca.

### Tujuan penerapan clean code :
* **Memudahkan Pemahaman**: Clean code adalah kode yang mudah dibaca dan dimengerti oleh pengembang lain atau diri sendiri. Ini membantu pengembang untuk lebih cepat memahami apa yang dilakukan kode tersebut, tujuan dari komponen-komponennya, dan bagaimana mereka saling berinteraksi.
* **Mendorong Kolaborasi**: Tim pengembangan biasanya terdiri dari beberapa anggota. Kode bersih membuatnya lebih mudah untuk berkolaborasi, karena setiap anggota tim dapat lebih cepat memahami dan berkontribusi pada proyek.
* **Meningkatkan Pemeliharaan**: Pengembangan perangkat lunak tidak hanya tentang menulis kode baru, tetapi juga tentang pemeliharaan kode yang sudah ada. Clean code memudahkan pemeliharaan karena kode yang bersih lebih mudah untuk dimodifikasi tanpa merusak fungsionalitas lain.
* **Mempercepat Pengembangan**: Clean code dapat menghemat waktu dalam jangka panjang. Kode yang sulit dipahami atau tidak terstruktur akan memperlambat pengembangan karena pengembang harus menghabiskan lebih banyak waktu mencari tahu cara kerja kode yang ada.

### Karakteristik clean code
* **Kode Bersih dan Mudah Dibaca**: Clean code harus mudah dibaca dengan formatasi yang konsisten, penggunaan nama variabel yang deskriptif, dan struktur yang terorganisasi dengan baik. Ini memungkinkan pengembang lain untuk memahami niat dan fungsionalitas kode tanpa kesulitan.

* **Menggunakan Nama Variabel yang Deskriptif**: Nama variabel, fungsi, kelas, dan metode harus memadai menjelaskan tujuan dan peran mereka dalam kode. Penggunaan nama yang singkat atau ambigu harus dihindari.

* **Fungsi dan Metode yang Kecil**: Fungsi dan metode harus berukuran kecil, menjalankan satu tugas yang jelas, dan mengikuti prinsip "Single Responsibility Principle" (SRP). Ini membuat kode lebih mudah diuji, dipahami, dan dikelola.

* **Penghindaran Kode Duplikat**: Kode yang berlebihan atau duplikat harus dihindari. Penggunaan fungsi atau kelas yang dapat digunakan kembali adalah cara yang baik untuk mengurangi duplikasi kode.

* **Komentar yang Dibutuhkan**: Komentar harus digunakan hanya jika diperlukan untuk menjelaskan kode yang kompleks atau algoritma yang tidak jelas. Idealnya, kode harus begitu jelas sehingga tidak memerlukan komentar.

* **Tidak Ada Kode Mati**: Kode mati atau kode yang tidak digunakan harus dihapus. Ini menjaga kode tetap bersih dan memudahkan pemeliharaan.

* **Menggunakan Konvensi Penamaan yang Konsisten**: Mengikuti konvensi penamaan yang konsisten untuk variabel, fungsi, dan kelas membantu dalam pemahaman dan pemeliharaan kode. Contohnya, penggunaan camelCase atau snake_case untuk penamaan variabel.

### Prinsip clean code
* **KISS (Keep It So Simple)**: Kode seharusnya sederhana dan tidak lebih kompleks dari yang diperlukan. Hindari over-engineering dan pertimbangkan solusi yang paling sederhana terlebih dahulu.
* **DRY (Don't Repeat Yourself)**: Hindari duplikasi kode. Setiap informasi seharusnya hanya memiliki satu sumber kebenaran di dalam kode.
