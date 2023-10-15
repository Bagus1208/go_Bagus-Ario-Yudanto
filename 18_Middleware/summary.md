# Summary

1. Middleware adalah sebuah kode yang dijalankan sebelum/setelah proses request/respone. Cara kerja middleware seperti penjaga di pintu masuk mall yang mengecek setiap pengunjung sebelum masuk ke dalam mall.

<br>

2. Ada dua tipe setup di middleware echo yaitu :
   * echo#pre()    : middleware akan dijalankan sebelum router memproses request
   * echo#user()   : middleware akan dijalankan setelah router memproses request dan mempunyai akses penuh ke API

<br>

3. Beberapa middleware yang sering digunakan yaitu :
   * CORS middleware, untuk memungkinkan API diakses dengan alamat/domain yang berbeda
   * Rate limiter middleware, untuk membatasi jumlah request ke server
   * Logging middleware, untuk mencatat informasi http request 
   * Authentication middleware (basic auth dan JWT), untuk mengidentifikasi user yang melalukan request


# Middleware
Middleware adalah entitas yang terkait dalam pemrosesan permintaan/respons server middleware, izinkan kami untuk melaksanakan berbagai fungsionalitas sekitar permintaan http masuk ke server anda dan respons keluar

### Cara kerja middleware
* **Penyisipan Middleware**: Middleware ditambahkan ke rantai pemrosesan permintaan (request processing chain) oleh pengembang. Saat permintaan tiba, middleware dijalankan secara berurutan sesuai dengan urutan penambahannya.

* **Permintaan (Request) Masuk**: Ketika seorang klien mengirim permintaan HTTP, permintaan tersebut pertama-tama melewati middleware sebelum mencapai tujuan akhir, seperti pengendali (controller) dalam aplikasi API.

* **Pemrosesan Middleware**: Middleware melakukan pemrosesan sesuai dengan fungsi atau tugas yang telah ditetapkan. Ini dapat mencakup:
  * **Otentikasi**: Middleware otentikasi dapat memeriksa identitas klien dan memastikan bahwa pengguna memiliki akses yang sesuai.
  Otorisasi: Middleware otorisasi memeriksa hak akses pengguna terhadap sumber daya atau tindakan tertentu.
  * **Logging**: Middleware logging dapat mencatat informasi seperti waktu permintaan, alamat IP klien, dan sebagainya.
  * **Kompressi**: Middleware kompresi dapat mengompres respons sebelum mengirimkannya ke klien untuk menghemat bandwidth.
  * **Manajemen Sesi**: Middleware sesi dapat mengelola data sesi pengguna, seperti token, cookie, atau state sesi.
* **Pengiriman Ke Tujuan Akhir**: Setelah pemrosesan middleware selesai, permintaan dapat dialihkan ke tujuan akhir yang sesuai, seperti pengendali (controller) dalam aplikasi API.

* **Generasi Respons**: Setelah pengendali (controller) menyelesaikan pemrosesan dan menghasilkan data respons, respons tersebut kembali melewati middleware dalam rantai respons (response chain) sebelum dikirimkan ke klien.

* **Pemrosesan Middleware Respons**: Middleware respons dapat digunakan untuk mengubah respons, menambahkan header, mengekstrak data, dan melakukan pemrosesan respons lainnya.

* **Pengiriman Respons ke Klien**: Akhirnya, respons dikirimkan ke klien.

## Middleware di Echo
* Basic Auth
* Body Dump
* Body Limit
* CORS
* CSRF
* Casbin Auth
* Gzip
* JWT
* Key Auth
* Logger
* Method Override
* Proxy
* Recover
* Redirect
* Request ID
* Rewrite
* Secure
* Session
* Static
* Trailing Slash

## HTTPS Redirect
```
e := echo. New()
e.Pre(middleware. HTTPSRedirect())
```

## CORS middleware
* CORS adalah singkatan dari Cross Origin Resource Sharing. 
* Dapat berbagi sumber daya dengan asal/domain berbeda. 
* Biasanya digunakan pada aplikasi web yang mengakses API dengan domain berbeda.

### Konfigurasi umum CORS
* **Access-Control-Allow-Origin**: Untuk menentukan domain / origin yang dapat mengirim request ke server.

* **Access-Control-Allow-Headers**: Untuk menentukan request header yang dapat digunakan ketika mengirim request ke server.

* **Access-Control-Allow-Methods**: Untuk menentukan HTTP Method yang dapat digunakan ketika mengirim request ke server.

* **Access-Control-Max-Age**: Untuk menentukan durasi maksimum preflight request yang dapat dilakukan caching.

### Implementasi CORS
```
package main

import (
   "net/http"
   "github.com/labstack/echo/v4"
   "github.com/labstack/echo/v4/middleware"
)

func main() {
   e := echo. New()

   // implement CORS middleware
   e.Use(middleware.CORS())

   e.GET("/", func(c echo.Context) error {
      return c.String(http.StatusOK, "hello")
   })

   e.Logger.Fatal(e.Start(":1323"))
}
```

### Implementasi CORS dengan konfigurasi
```
func main() {
   e := echo.New()

   e.Use(middleware.Logger())

   // implement CORS middleware with custom configurations
   e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      AllowOrigins: []string{"http://localhost:8080", "http://localhost: 8081"},
      AllowHeaders: []string{echo. HeaderOrigin, echo. HeaderContentType, echo. HeaderAccept},
      AllowMethods: []string{http.MethodGet, http.MethodPost},
   }))

   e.GET("/", func(c echo.Context) error {
      return c.String(http.StatusOK, "hello")
   })

   e.Logger.Fatal(e.Start(":1323"))
}
```

## Rate limiter
* Batasi jumlah permintaan ke server. 
* Cegah permintaan pengambilan berlebih untuk memastikan kinerja server tetap terjaga. 
* Cegah pelanggaran keamanan seperti DDos, Brute Force Attack.

### Implementasi rate limiter
```
package main

import (
   "net/http"
   "github.com/labstack/echo/v4"
   "github.com/labstack/echo/v4/middleware"
)

func main() {
   e := echo. New()

   // implement rate limiter middleware
   // limit to 20 requests/second
   e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

   e.GET("/", func(c echo.Context) error {
      return c.String(http.StatusOK, "hello")
   })

   e.Logger.Fatal(e.Start(":1323"))
}
```

### Implementasi rate limiter dengan konfigurasi
```
func main( ) {
   e := echo.New()

   config := middleware. RateLimiterMemoryStoreConfig{
      Rate: 10,                     // rate limit to 10 request/sec
      Burst: 30,                    // burst up to 30 requests
      ExpiresIn: 3 * time.Minute,   // time expiration
   }

   // implement rate limiter middleware with configuration
   e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(config)))

   e.GET("/", func(c echo. Context) error {
      return c.String(http.StatusOK, "hello")
   })
   e.Logger.Fatal(e.Start(":1323"))
}
```

## Logger middleware
* Mencatat informasi permintaan HTTP. 
* Sebagai footprint. 
* Sumber data untuk analitik.

### Implementasi logger
```
package middlewares

import (
   github.com/labstack/echo/v4"
   "github.com/labstack/echo/v4/middleware"
)

func LogMiddlewares(e *echo. Echo) {
   e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
      Format: [${time_rfc3339} ] ${status} ${method}${host}${path} ${latency_human} + "\n",
   }))
}
```
```
package main

import(
   "github.com/iswanulumam/go-training-restful/routes"
   m "github.com/iswanulumam/go-training-restful/middleware
)

func main( ) {
   e := routes. New()

   // implement middleware logger
   m. LogMiddlewares(e)

   // start the server, and log if it fails
   e.Logger.Fatal(e.Start(":8000"))
}
```

## Authentication middleware
* **Basic Authentication Middleware**: Middleware ini memeriksa otentikasi pengguna dengan menggunakan skema otentikasi dasar (Basic Authentication). Klien harus mengirimkan header Authorization dengan informasi otentikasi (nama pengguna dan kata sandi) yang dienkripsi dalam format base64. Jika informasi otentikasi valid, permintaan diteruskan; jika tidak, klien diminta otentikasi ulang.

* **JWT (JSON Web Token) Authentication Middleware**: Middleware ini memeriksa token JWT yang dikirim oleh klien dalam header permintaan. Token JWT adalah tanda tangan digital yang dapat digunakan untuk mengautentikasi pengguna. Middleware ini memeriksa keaslian token, waktu kedaluwarsa, dan hak akses pengguna.