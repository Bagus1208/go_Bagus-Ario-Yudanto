# Docker
Docker adalah platform untuk mengembangkan, mengemas, dan menjalankan aplikasi dan dependensi dalam lingkungan yang disebut kontainer. kontainer adalah unit perangkat lunak yang mandiri, yang mencakup semua yang diperlukan untuk menjalankan aplikasi, termasuk kode, runtime, pustaka, dan dependensi.

## Istilah dalam docker
* **Dockerfile** : file teks yang berisi instruksi untuk membangun Docker image. Dockerfile digunakan untuk menggambarkan langkah-langkah yang diperlukan untuk membuat image yang akan digunakan sebagai dasar kontainer.
* **Docker Image** : tamplate untuk membuat kontainer. Image berisi instruksi tentang apa yang harus dimasukkan dalam kontainer, seperti sistem operasi dasar, aplikasi, dan konfigurasi. Image digunakan untuk membuat instance kontainer yang dapat dijalankan.
* **Docker Hub** : repositori pusat untuk Docker images. Anda dapat menemukan dan mengunduh image yang telah dibuat sebelumnya dari Docker Hub, atau Anda dapat mengunggah image buatan sendiri ke sana untuk berbagi dengan komunitas.
* **Docker Compose** : alat untuk mendefinisikan dan menjalankan aplikasi multi-container dengan menggunakan satu file konfigurasi. Ini memudahkan pengembang untuk mengatur dan menghubungkan kontainer yang berbeda dalam satu aplikasi.
* **Docker Registry** : wadah untuk menyimpan Docker image. Docker image akan memberi reaksi sesuai perintah yang diberikan. Misalnya, saat diberi perintah docker push, docker image akan didorong atau dibagikan ke registry Docker Hub.

## Fungsi docker
* Mempermudah pengembangan aplikasi
* Menyederhanakan konfigurasi
* Memudahkan pengembangan
* Bisa digunakan untuk debugging
* Mendukung multitenancy

## Sintak didalam dockerfile
* **FROM** : mendapatkan image dari docker registry
* **RUN** : eksekusi perintah bash ketika membuat kontainer
* **ENV** : set variabel didalam kontainer
* **ADD** : menyalin file dengan beberapa proses lainnya
* **COPY** : menyalin file
* **WORKDIR** : set direktori kerja file
* **ENTRYPOINT** : eksekusi perintah ketika selesai membuat kontainer
* **CMD** : eksekusi perintah tetapi bisa ditimpa

### Contoh dockerfile
```
FROM golang:1.16-alpine
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -o app .
CMD ["/app/app"]
```

## Perintah-perintah di docker
### Membuat docker image
```
docker build -t <nama_image>
```

### Login ke docker registry
```
docker login --username=yourusername
```

### Tag dan push docker image
```
docker tag 566hhfg5446g ian18ishar/flask-tutorial:1.0

docker push ian18ishar/flask-tutorial
```

### Pull dan menjalankan kontainer
```
docker pull ian18ishar/flask-tutorial:1.0

docker run -d -p 5000:5000 --name flaskdemo ian18ishar/flask-tutorial:1.0
```