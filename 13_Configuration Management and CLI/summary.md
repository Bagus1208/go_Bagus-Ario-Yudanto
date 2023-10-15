# Summary

1. Command Line Interface adalah interface berbasis text yang dapat memberi perintah ke komputer untuk mengerjakan tugas dengan mengetik perintah tertentu.

<br>

2. Kelebihan CLI yaitu :
   * Kontrol secara terperinci ke sistem operasi atau aplikasi
   * Manajemen lebih cepat
   * Memiliki kemampuan menyimpan kumpulan perintah untuk menjalankan tugas secara otomatis
   * Efisien dalam penggunaan hardware karena tidak memakan banyak resource

<br>

3. Shell adalah program yang menjembatani antara user dengan kernel. Shell script bahasa program yang dikompilasi berdasarkan perintah shell.

# Configuration Management and CLI
## CLI (Command Line Interface)
CLI adalah singkatan dari "Command Line Interface." Ini adalah antarmuka pengguna yang memungkinkan pengguna untuk berinteraksi dengan komputer atau sistem operasi melalui baris perintah teks.

### Alasan menggunakan CLI
* Kontrol granular dari suatu OS atau aplikasi
* Manajemen lebih cepat dari sejumlah besar sistem operasi
* Kemampuan untuk menyimpan skrip untuk mengotomatiskan tugas-tugas rutin
* Menghubungkan pengetahuan untuk membantu pemecahan masalah, seperti masalah koneksi jaringan

## Unix Shell
Unix shell adalah program berbasis baris perintah yang digunakan untuk berinteraksi dengan sistem operasi Unix atau Unix-like (seperti Linux). Shell bertindak sebagai antarmuka antara pengguna dan kernel sistem operasi, memungkinkan pengguna untuk menjalankan perintah dan mengelola sistem menggunakan perintah-perintah teks.

Contoh unix shell:
```
me@linuxbox:~$
```
* (me) -> your username
* (linuxbox) -> your hostname
* (~) -> your current path (home)
* ($) -> Shell for normal username
* (#)-> Shell for root user

### Normal user vs Root user
* Normal User = Allowed to manupulating /home/username dir only
* Root User = Allowed to manupulating /dir (all directory)
* Normal User + Sudoers = Allowed to act as root temporary

## Unix Shell Command
### Directory
  * pwd
  * ls
  * mkdir 
  * cd
  * rm
  * cp
  * mv
  * ln

### Files
  * create : touch
  * view : head, cat, tail,less
  * editor : vim, nano
  * permission : chown, chmod
  * Different : diff

### Network
* ping
* ssh
* netstat
* nmap
* ip addr (ifconfig)
* wget
* curl
* telnet
* netcat

### Utility
* man
* env
* echo
* date
* which
* watch
* sudo
* history
* grep
* locate