# Summary

1. Version Control System (VCS) adalah sistem yang merekam perubahan pada sebuah file dari waktu ke waktu sehingga kita dapat melihat versi sebelumnya jika diinginkan.
   
<br>

2. Version Control System dibagi menjadi tiga yaitu :
   * Single User : 
      - SCCS 
      - RCS
   * Centralized : 
      - CVS, Perforce 
      - Subversion
      - Microsoft Team Foundation Server
   * Distributed :
      - Git  
      - Mercurial
      - Bazaar

<br>

3. Untuk mengoptimalkan workflow dalam bekerja sama menggunakan github atau gitlab ada beberapa cara yaitu :
   * Jangan melakukan perbaikan/penambahan fitur di master branch. Buat development branch yang mengacu pada master branch
   * Hindari perbaikan/penambahan fitur di development branch. Buat branch baru untuk setiap perubahan/penambahan fitur
   * Terapkan fitur yang sudah diperbaiki/ditambahkan ke development branch
   * Terapkan development branch ke master branch jika fitur yang diperbaiki/ditambahkan sudah berjalan dengan baik dan tidak ada masalah 

# Version Control and Branch Management (Git)
## Version Control System
Version control system adalah sistem yang merekam perubahan pada sebuah file dari waktu ke waktu sehingga kita dapat melihat versi sebelumnya jika diinginkan.

### Tipe version control system
1. Local version control
2. Centralize version control
3. Distributed version control

## Git
Git adalah salah satu version contorl system populer yang digunakan para developer untuk mengembangkan software secara bersamaan

### Konfigurasi Git
* Konfigurasi user name dan user email
  ```
  git config --global user.name "your name"
  git config --global user.email "youremail"
  ```
* Konfigurasi menjadikan VS Code sebagai default editor dan default diff tool
  ```
  git config --global core.editor "code --wait"
  git config --global diff.tool "default-difftool"
  git config --global difftool.default-difftool.cmd "code --wait -diff $LOCAL $REMOTE"
  ```
* Melihat seluruh konfigurasi
  ```
  git config --list --show-origin
  ```
## Repository
Repository adalah istilah yang digunakan untuk direktori yang menggunakan Git. 
### Cara membuat repository 
```
git init
```

### Tiga state di dalam Git
1. Modified, kita mengubah(menambah, mengedit, menghapus) file, tetapi belum disimpan secara permanen ke repository
2. Staged, kita menandai modifikasi yang kita lakukan terhadap file akan disimpan secara permanen ke repository
3. Committed, file sudah tersimpan di repository

### Tiga Section di dalam git
Tiga state sebelumnya di dalam Git dilakukan di section berbeda-beda, yaitu :
1. Working, saat kita melakukan perubahan file
2. Staging area, saat menandai modifikasi file yang akan disimpan ke repository dan semua informasi perubahan file akan disimpan
3. Repository, tempat semua file dan database riwayat versi file disimpan

## Branch Management
Branch artinya kita membuat timeline baru yang berbeda dengan timeline utama. Timeline utama atau branch utama disebut dengan master atau main. Saat kita membuat branch baru dan melakukan perubahan di branch tersebut, perubahan tidak akan merusak timeline branch utama. 

### Cara membuat branch baru
1. Tidak langsung pindah ke branch yang dibuat
   ```
   git branch featureA
   ```
2. Langsung pindah ke branch yang dibuat
   ```
   git checkout -b featureA
   ```

### Cara menggabungkan branch
1. Pindah ke branch dimana tempat melakukan penggabungan
2. Lakukan perintah
   ```
   git merge featureA
   ```

### Strategi branch management
1. [Gitflow Workflow](https://nvie.com/posts/a-successful-git-branching-model/)
2. [Trunk Based Development Workflow](https://trunkbaseddevelopment.com/)
3. [Forking Workflow](https://www.atlassian.com/git/tutorials/comparing-workflows/forking-workflow)




