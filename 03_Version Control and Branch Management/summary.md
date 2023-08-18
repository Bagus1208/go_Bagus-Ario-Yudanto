# Summary

1. Version Control System (VCS) adalah sistem yang merekam perubahan pada sebuah file dari waktu ke waktu sehingga kita dapat melihat versi sebelumnya jika diinginkan.
   
<br>

2. Version Control System dibagi menjadi tiga yaitu :
   1. Single User : 
      - SCCS 
      - RCS
   2. Centralized : 
      - CVS, Perforce 
      - Subversion
      - Microsoft Team Foundation Server
   3. Distributed :
      - Git  
      - Mercurial
      - Bazaar

<br>

3. Untuk mengoptimalkan workflow dalam bekerja sama menggunakan github atau gitlab ada beberapa cara yaitu :
   1. Jangan melakukan perbaikan/penambahan fitur di master branch. Buat development branch yang mengacu pada master branch
   2. Hindari perbaikan/penambahan fitur di development branch. Buat branch baru untuk setiap perubahan/penambahan fitur
   3. Terapkan fitur yang sudah diperbaiki/ditambahkan ke development branch
   4. Terapkan development branch ke master branch jika fitur yang diperbaiki/ditambahkan sudah berjalan dengan baik dan tidak ada masalah 