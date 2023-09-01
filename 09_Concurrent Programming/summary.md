# Summary 

1. Concurrent programming adalah program yang menjalankan beberapa tugas bergantian secara independen. Implementasi concurrent programming di golang adalah goroutine.

<br>

2. Perbedaan antara sequential, parallel, dan concurrent
   1. Sequential
      - sebelum tugas baru dijalankan, tugas yang sedang dijalankan harus selesai dahulu.
   2. Parallel
      - beberapa tugas dapat dijalankan bersamaan.
   3. Concurrent
      - Beberapa tugas dijalankan secara bergantian hingga selesai.

<br>

3. Pada concurrent programming terkadang muncul masalah yaitu race condition dimana 2 thread atau lebih mengakses memory yang sama dalam waktu yang sama. Untuk mengatasi itu ada beberapa cara yaitu dengan menggunakan :
   1. waitgroup
   2. channel blocking
   3. mutex