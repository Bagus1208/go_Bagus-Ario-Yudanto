# Summary

1. Echo merupakan framework web golang yang dibuat oleh labstack yang dioptimalkan dengan alokasi memori dinamis nol yang memprioritaskan routes dengan cerdas

<br>

2. Echo merupakan framework yang minimalis sehingga memiliki performa yang baik. Dan echo memberikan kontrol penuh kepada developer dalam merancang aplikasi web.

<br>

3. Keunggulan echo antara lain :
   * Mengoptimalkan router
   * Dapat diperluas
   * Dilengkapi dengan banyak built-in middleware
   * Berisi API untuk mengirim berbagai respon http termasuk JSON, XML, HTML, file, attachment, inline, stream atau blob
   * Mendukung data binding untuk http request berupa JSON, XML, atau form-data

# Intro Echo Golang
Echo merupakan framework web golang yang dibuat oleh labstack yang dioptimalkan dengan alokasi memori dinamis nol yang memprioritaskan routes dengan cerdas

## Alasan menggunakan Echo
* **Kinerja Tinggi**: Echo dirancang untuk kinerja tinggi. Dengan pendekatan yang sangat efisien dan ringan, itu dapat menangani lalu lintas yang padat dengan cepat.

* **Sederhana dan Mudah Dipelajari**: Echo memiliki antarmuka yang sederhana dan mudah dipahami. Ini membuatnya cocok untuk pemula dan pengembang berpengalaman. Anda tidak perlu belajar banyak konsep rumit untuk mulai menggunakan Echo.

* **Rute yang Kuat**: Echo menyediakan rute yang kuat yang memungkinkan Anda untuk dengan mudah menangani permintaan HTTP dan menghubungkannya dengan fungsi yang sesuai dalam aplikasi Anda.

* **Penanganan Middleware**: Echo mendukung middleware yang memungkinkan Anda untuk menambahkan fungsionalitas tambahan seperti otentikasi, log, kompresi, dan sebagainya ke aplikasi Anda dengan mudah.

* **Template Mudah**: Echo memiliki dukungan untuk template HTML, yang memungkinkan Anda untuk dengan mudah merender halaman web dinamis.

* **Dokumentasi yang Baik**: Echo memiliki dokumentasi yang baik dan komunitas yang aktif. Ini membuatnya lebih mudah untuk memahami dan menyelesaikan masalah.

* **Ekosistem yang Kuat**: Seiring dengan perkembangan Go sebagai bahasa pemrograman, ekosistem Echo juga berkembang. Ada banyak pustaka tambahan dan alat yang dapat digunakan bersama dengan Echo.

* **RESTful API**: Echo memiliki dukungan bawaan untuk membangun RESTful API, yang sangat berguna jika Anda ingin membuat layanan web yang menyediakan dan mengelola data.

* **Banyak Fitur Opsional**: Echo menyediakan banyak fitur opsional seperti dukungan Gzip, middlewares pihak ketiga, WebSocket, dan sebagainya yang dapat Anda gunakan sesuai kebutuhan.

## Setup Echo
```
$ go get github.com/labstack/echo/v4
```

## Penggunaan Echo
```
package main

import (
   "net/http"
   "github.com/labstack/echo"
)

func main() {
   // create a new echo instance
   e := echo. New()
   // Route / to handler function
   e.GET("/", HelloController)
   // start the server, and log if it fails
   e.Start(":8000")
} 

// handler - Simple handler to make sure environment is setup
func HelloController(c echo.Context) error {
   // return the string "Hello World" as the response body
   // with an http. StatusOK (200) status
   return c.String(http.StatusOK, "Hello World")
} 
```

