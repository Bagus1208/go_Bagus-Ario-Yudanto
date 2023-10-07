# Compute service
Compute service adalah adalah istilah umum yang digunakan dalam konteks komputasi awan (cloud computing) untuk merujuk kepada layanan yang menyediakan daya komputasi seperti server virtual, pemrosesan, dan sumber daya komputasi lainnya. Compute service memungkinkan pengguna untuk menjalankan aplikasi, menyimpan data, dan melakukan pemrosesan di cloud tanpa harus memiliki infrastruktur fisik sendiri.

### Contoh compute service 
* Amazon EC2
* GCP
* Digital Ocean

## System and Software Deployment
Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. Penyebaran dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web & api ke server sedangkan aplikasi mobile ke Playstore/Appstore.

## Strategi Deployment
### Big-Bang Deployment Strategy/Replace Deployment Strategy
Kelebihan :
* Mudah diimplementasikan. Cara klasik, tinggal di replace
* Perubahan kepada sistem langsung 100% secara instan

Kekurangan :
* Terlalu beesiko
* Rata-rata downtime cukup lama
  
### Rollout Deployment Strategy
Kelebihan :
* Lebih aman dan less downtime dari versi sebelumnya

Kekurangan :
* Akan ada 2 versi aplikasi berjalan bersamaan sampai semua server terdeploy 
* Untuk deployment dan rollback lebih lama dari Big-Bang
* Tidak ada kontrol request
  
### Blue/Green Deployment Strategy
Kelebihan :
* Perubahan sangat cepat, sekali switch service langsung berubah 100%
* Tidak ada issue beda versi aplikasi

Kekurangan :
* Resource yang dibutuhkan lebih banyak karena perlu menyediakan service yang serupa environmentnya dengan service yang berjalan
* Testing harus benar-benar sangat diprioritaskan 
  
### Canary Deployment Strategy
Kelebihan :
* Cukup aman
* Mudah untuk rollback jika terjadi error/bug tanpa berimbas ke semua user

Kekurangan :
* Proses deployment lebih lambat dari Blue/Green deployment