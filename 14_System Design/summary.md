# Summary

1. Design software yang umum digunakan :
   * Flowchart
   * Use case
   * Entity Relational Diagram (ERD)
   * High Level Architecture (HLA)

<br>

2. Kunci karakteristik sistem terdistribusi :
   * Scalability
   * Reliability
   * Availability
   * Efficiency
   * Serviceability or manageability

<br>

3. Load balancing berguna untuk membagi beban pekerjaan ke beberapa server untuk meningkatkan respon dan ketersediaan aplikasi. Load balancing dapat ditempatkan di :
   * antara user dan web server
   * antara web server dan internal platform seperti server aplikasi atau server cache
   * antara internal platform dan database

# System Design
System design, atau desain sistem, adalah proses perencanaan dan pengembangan struktur, komponen, dan fitur-fitur yang akan ada dalam sebuah sistem komputer atau perangkat lunak. Tujuan utama dari system design adalah untuk merancang sistem yang memenuhi kebutuhan fungsional, kinerja, keamanan, dan keandalan yang ditentukan dalam spesifikasi awal.

### Desain perangkat lunak yang biasa digunakan
* Flowchart
* Use Case
* ERD (Entity Relational Database)
* HLA (High Level Architecture)

## karakteristik utama dari sistem terdistribusi
### Scalability
Scalability adalah kemampuan dari sebuah sistem, proses, atau jaringan untuk tumbuh dan mengelola peningkatan permintaan. Setiap sistem terdistribusi yang dapat terus berkembang untuk mendukung peningkatan jumlah pekerjaan dianggap dapat ditingkatkan kapasitasnya. Sebuah sistem mungkin harus meningkatkan kapasitasnya karena banyak alasan seperti peningkatan volume data atau peningkatan jumlah pekerjaan, misalnya, jumlah transaksi. Sebuah sistem yang dapat ditingkatkan kapasitasnya ingin mencapai peningkatan ini tanpa kehilangan kinerja.

### Reability
Secara definisi, reability adalah probabilitas sebuah sistem akan gagal dalam periode tertentu. Dalam istilah sederhana, sebuah sistem terdistribusi dianggap dapat diandalkan jika terus memberikan layanannya bahkan ketika satu atau beberapa komponen perangkat lunak atau perangkat kerasnya gagal. Reliabilitas mewakili salah satu karakteristik utama dari setiap sistem terdistribusi, karena dalam sistem seperti itu, mesin yang gagal selalu dapat digantikan oleh mesin yang sehat lainnya, memastikan penyelesaian tugas yang diminta.

### Availability
Secara definisi, availability adalah waktu di mana sistem tetap beroperasi untuk melakukan fungsi yang dibutuhkan dalam periode tertentu. Ini adalah ukuran sederhana dari persentase waktu bahwa sistem, layanan, atau mesin tetap beroperasi dalam kondisi normal. Contoh: sebuah pesawat yang dapat terbang selama banyak jam dalam sebulan tanpa banyak waktu henti dapat dikatakan memiliki ketersediaan yang tinggi.

### Efficiency
Untuk memahami bagaimana mengukur efisiensi dari sebuah sistem terdistribusi, mari kita asumsikan kita memiliki operasi yang berjalan secara terdistribusi dan menghasilkan serangkaian item sebagai hasilnya. Dua ukuran standar dari efisiensinya adalah waktu respon (atau latensi) yang menunjukkan penundaan untuk mendapatkan item pertama, dan throughput (atau bandwidth), yang menunjukkan jumlah item yang dikirim dalam satuan waktu tertentu (misalnya, satu detik).

### Serviceability & Manageability
Serviceability dan Manageability adalah kemudahan dan kecepatan dengan mana sistem dapat diperbaiki atau dipelihara; jika waktu untuk memperbaiki sistem yang gagal meningkat, maka ketersediaan akan menurun. Hal-hal yang perlu dipertimbangkan untuk manajemen adalah kemudahan dalam mendiagnosis dan memahami masalah ketika mereka terjadi, kemudahan dalam membuat pembaruan atau modifikasi, dan seberapa sederhana sistem dioperasikan (misalnya, apakah sistem secara rutin beroperasi tanpa kegagalan atau pengecualian?).
   
## Load Balancing
Load Balancer (LB) adalah komponen penting lainnya dari setiap sistem terdistribusi. Ini membantu menyebarkan lalu lintas di sekelompok server untuk meningkatkan responsivitas dan ketersediaan aplikasi, situs web, atau basis data. LB juga melacak status semua sumber daya saat mendistribusikan permintaan. Jika server tidak tersedia untuk menerima permintaan baru atau tidak merespons atau memiliki tingkat kesalahan yang meningkat, LB akan berhenti mengirimkan lalu lintas ke server tersebut.

Untuk memanfaatkan skalabilitas dan redundansi secara penuh, kita bisa mencoba menyeimbangkan beban di setiap lapisan sistem. Kita bisa menambahkan Load Balancer di tiga tempat:

* Antara pengguna dan server web
* Antara server web dan lapisan platform internal, seperti server aplikasi atau server cache
* Antara lapisan platform internal dan database.

## Monolithic dan Microservices
### Monolithic
Aplikasi monolitik memiliki basis kode tunggal dengan beberapa modul. Modul dibagi berdasarkan fitur bisnis atau fitur teknis. Aplikasi ini memiliki sistem build tunggal yang membangun seluruh aplikasi dan/atau dependensinya. Aplikasi ini juga memiliki biner yang dapat dieksekusi atau dikerahkan secara tunggal.
### Microservices
Microservices adalah layanan yang dapat dikerahkan secara independen dan dimodelkan berdasarkan domain bisnis. Mereka berkomunikasi satu sama lain melalui jaringan, dan sebagai pilihan arsitektur menawarkan banyak opsi untuk memecahkan masalah yang mungkin Anda hadapi. Dengan demikian, arsitektur microservice didasarkan pada beberapa microservice yang saling bekerja sama.

## SQL & NoSQL
* Basis data relasional memiliki struktur dan skema yang telah ditentukan sebelumnya seperti buku telepon yang menyimpan nomor telepon dan alamat.
* Basis data non-relasional tidak terstruktur, dan memiliki skema dinamis seperti folder yang menyimpan segala sesuatu mulai dari alamat dan nomor telepon seseorang hingga ‘suka’ Facebook dan preferensi belanja online mereka.

## Caching
Cache digunakan ketika data yang baru saja diminta kemungkinan akan diminta lagi. Cache digunakan di hampir setiap lapisan komputasi: perangkat keras, sistem operasi, browser web, aplikasi web, dan lainnya. Cache seperti memori jangka pendek: memiliki ruang yang terbatas, tetapi biasanya lebih cepat daripada sumber data asli dan berisi item yang paling baru diakses. Cache dapat ada di semua tingkat arsitektur, tetapi sering ditemukan di tingkat terdekat ke front end di mana mereka diimplementasikan untuk mengembalikan data dengan cepat tanpa membebani tingkat hulu.

## Database Replication
Redundansi adalah duplikasi komponen atau fungsi kritis dari suatu sistem dengan tujuan meningkatkan keandalan sistem, biasanya dalam bentuk cadangan atau fail-safe, atau untuk meningkatkan kinerja sistem sebenarnya. Misalnya, jika hanya ada satu salinan file yang disimpan pada satu server, maka kehilangan server tersebut berarti kehilangan file tersebut. Karena kehilangan data jarang merupakan hal yang baik, kita dapat membuat salinan duplikat untuk memecahkan masalah tersebut.

## Database Indexing
Indexing adalah cara untuk mengoptimalkan kinerja database dengan meminimalkan jumlah akses disk yang diperlukan saat query diproses. Ini adalah teknik struktur data yang digunakan untuk dengan cepat menemukan dan mengakses data dalam database.