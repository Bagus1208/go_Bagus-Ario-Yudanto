# Summary

1. API (Aplication Programming Interface) adalah kumpulan fungsi dan prosefur yang memungkinkan pembuatan aplikasi mengakses fitur atau data dari suatu sistem lainnya. 

<br>

2. REST (Representational State Transfer) adalah serangkaian prinsip dan arsitektur software. RESTful API adalah implementasi yang mematuhi semua batasan dan pedoman dari REST.

<br>

3. HTTP response code diklasifikasikan menjadi 5 yaitu :
   * 1xx untuk respon informasi. Contoh 100 Continue, 101 Switching Protocols
   * 2xx untuk sukses. Contoh 200 OK, 201 Created
   * 3xx untuk pengalihan. Contoh 301 Moved Permanently, 303 See Other
   * 4xx untuk error pada client. Contoh 400 Bad Request, 401 Unauthorized
   * 5xx untuk server error. Contoh 500 Internal Server Error, 502 Bad Gateway 

# Intro RESTful API
## API 
API (Aplication Programming Interface) adalah kumpulan fungsi dan prosefur yang memungkinkan pembuatan aplikasi mengakses fitur atau data dari suatu sistem lainnya.

### Cara kerja API
* **Permintaan (Request)**: Sebuah aplikasi mengirim permintaan kepada API dengan permintaan tertentu, yang mencakup metode (seperti GET, POST, atau PUT) dan data yang diperlukan.

* **Pemrosesan (Processing)**: API memproses permintaan tersebut, melakukan tindakan yang sesuai, seperti mengambil data dari database, memproses data, atau menjalankan operasi lain.

* **Respon (Response)**: Setelah pemrosesan selesai, API menghasilkan respon yang berisi hasil dari permintaan. Respon ini dapat berupa data dalam format yang sesuai, seperti JSON atau XML.

* **Pengiriman Respon (Response Delivery)**: Respon dikirim kembali ke aplikasi yang melakukan permintaan, biasanya melalui jaringan atau internet.

* **Pengolahan Respon (Response Processing)**: Aplikasi yang melakukan permintaan menerima respon dan melakukan pemrosesan lebih lanjut sesuai dengan kebutuhan, seperti menampilkan data kepada pengguna.

## REST 
REST (Representational State Transfer) adalah gaya arsitektur yang digunakan dalam pengembangan aplikasi web dan layanan web. Ini didasarkan pada sejumlah prinsip desain yang mempromosikan skalabilitas, fleksibilitas, dan kesederhanaan dalam komunikasi antara berbagai komponen perangkat lunak.

### Cara kerja REST
* **Sumber Daya (Resource)**: REST berfokus pada sumber daya, yang dapat menjadi objek atau data tertentu. Setiap sumber daya memiliki identifikasi unik yang dikenal sebagai URI (Uniform Resource Identifier).

* **Operasi pada Sumber Daya**: Untuk berinteraksi dengan sumber daya, klien (misalnya, aplikasi web atau perangkat mobile) mengirim permintaan HTTP ke URI sumber daya yang tepat. Permintaan HTTP dapat berupa GET (mengambil data), POST (mengirim data baru), PUT (mengganti data yang ada), atau DELETE (menghapus data).

* **Representasi**: Sumber daya dapat memiliki beberapa representasi, seperti HTML, XML, JSON, atau teks biasa. Klien meminta atau mengirim data dalam format yang diinginkan melalui header permintaan HTTP.

* **Stateless**: REST adalah arsitektur stateless, yang berarti setiap permintaan dari klien ke server harus mengandung semua informasi yang diperlukan. Server tidak menyimpan informasi tentang status klien sebelumnya.

* **Sistem Terdistribusi**: REST mendukung sistem terdistribusi, yang berarti sumber daya dan layanan dapat tersebar di banyak server yang berbeda.

* **Penggunaan HTTP**: REST menggunakan protokol HTTP sebagai media komunikasi utamanya. Ini memanfaatkan metode HTTP, seperti GET, POST, PUT, dan DELETE, serta kode status HTTP, seperti 200 (OK), 201 (Created), dan lainnya.

* **HATEOAS (Hypermedia as the Engine of Application State)**: Prinsip ini menekankan bahwa API REST harus menyertakan tautan (link) ke sumber daya lain, memungkinkan klien untuk menavigasi aplikasi dengan lebih dinamis.

## JSON
JSON adalah singkatan dari "JavaScript Object Notation." Ini adalah format data ringan dan terbaca oleh manusia yang digunakan untuk pertukaran data antara aplikasi.

Contoh JSON:
```
{
   "name": "Anam",
   "umur": 17,
   "single": true,
   "hobi": [
      "belajar",
      "mengaji"
   ],
   "alamat": {
      "rumah": "Bogor",
      "no": 4,
      "rt": "03",
      "rw": "11"
   }
}
```

## HTTP Response Code
* 1xx untuk respon informasi. Contoh 100 Continue, 101 Switching Protocols
* 2xx untuk sukses. Contoh 200 OK, 201 Created
* 3xx untuk pengalihan. Contoh 301 Moved Permanently, 303 See Other
* 4xx untuk error pada client. Contoh 400 Bad Request, 401 Unauthorized
* 5xx untuk server error. Contoh 500 Internal Server Error, 502 Bad Gateway 

## API tools
* Katalon
* Apache JMeter
* Postman
* SoapUI

## REST API Design Best Practices
* Gunakan kata benda daripada kata kerja
  * Gunakan `GET /books/123` daripada `GET /addbooks/123

  
* Penamaan menggunakan kata benda jamak 
  * Gunakan `GET /books/123` daripada `GET /book/123`
  
* Menggunakan sumber daya bersarang untuk memperlihatkan relasi atau hierarki
   ```
   /users <- user's list
   /users/123 <- specific user
   /users/123/orders <- orders list that belongs to a specific user
   /users/123/orders/0001 <- specific order of a specific user
   ```
* Format response JSON
* Filtering, Sorting, Paging, and Field
Selection
  * Filtering:
      ```
      GET /users?country=USA
      GET /users?creation_date=2019-11-11
      GET /users?creation_date=2019-11-11
      ```
  * Sorting:
      ```
      GET /users?sort=birthdate_date:asc
      GET /users?sort=birthdate_date:desc
      ```
  * Paging:
      ```
      GET /users?limit=100
      GET /users?offset=2
      ```

  * Field Selection:
      ```
      GET /users/123?fields=username,email
      (for one specific user)

      GET /users?fields=username,email
      (for a full list of users)
      ```

* Tangani Trailing Slash dengan Anggun
  * `POST: /cars` & `POST: /cars/`

* Versioning
   ```
   https://us6.api.mailchimp.com/3.0/ (major + minor version indication)

   https://api.stripe.com/v1/ (major version indication only)

   https://developer.github.com/v3/ (major version indication only)
   ```

* API documentation
  * Ingat bahwa dokumentasi API juga menunjukkan seberapa besar kepedulian perusahaan Anda

