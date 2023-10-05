# Clean dan Hexagonal Architecture  
Clean dan hexagonal architecture adalah arsitektur aplikasi dengan memisahkan setiap bagian kode sesuai dengan fungsinya yang bertujuan agar kode bersifat modular, mudah dikembangkan, mudah di perbaiki, dan mudah dalam melakukan pengujian.

### Tujuan clean dan hexagonal architecture
Bekerja dalam projek pemrograman sama dengan merencanakan dan membangun daerah perumahan. Jika kita tahu bahwa daerah tersebut akan berkembang di kemudian hari, kita perlu menyediakan tempat untuk pengembangan di masa depan. Tanpa itu, kita perlu menghancurkan paksa bangunan dan jalan untuk tempat bangunan baru

### Aturan dalam mendesain clean architecture
* Independent of framework, ini memungkinkan untuk menggunakan framework sebagai alat, daripada memiliki batasan dari framework tersebut.
* Testable, business rules dapat melakukan pengujian tanpa melibatkan module yang lain
* Independent of UI, tampilan dapat dengan mudah berubah, tanpa merubah fungsi dari sistem.
* Independent database, business rules tidak bergantung penuh ke database sehingga dapat mengganti database dengan mudah
* Independent of any external, business rules tidak perlu mengetahui sistem eksternal yang digunakan


### Manfaat menerapkan clean architecture
* struktur standar, sehingga mudah dalam membangun sebuah aplikasi
* Pengembangan yang lebih cepat dalam jangka panjang
* Ketergantungan mocking menjadi hal yang sepele di unit test
* pergantian yang mudah dari prototype ke solusi yang tepat(contoh pergantian in-memory storage ke SQL database)  

### Layer dalam clean architecture
* Entities layer, berisi business object yang mencerminkan konsep dari dikelola aplikasi
* Use case - domain layer, berisi business logic
* Controller - presentation layer, berisi tampilan data dan menangani interaksi pengguna
* Driver - data layer, mengelola data aplikasi, contoh mengambil data dari jaringan, mengelola data cache

Contoh penerapan clean architecture :
```
myapp/
├── feature/
│   └── users/
|       ├── controller/
│       │   └── controller.go
│       ├── repository/
│       │   └── repository.go
│       ├── usecase/ 
│       │   └── usecase.go 
|       └── entity.go
└── main.go
```

# Context Golang
Package context atau kita sebut saja context, adalah suatu package yang membawa dealine, cancellation signal, atau value lain pada request/permintaan API

## Implementasi context
Kita bisa membuat root dengan fungsi Background()

```
var ctx = context.Background()
```

Fungsi Background() akan mengembalikan root context dimana kita bisa memakainya untuk operasi lain

## Context with value
```
var ctx = context.Background()

ctx = context.WithValue(ctx, "key", "value")
```

## Context with cancellation
```
var parent = context.Background()

ctx = context.WithCancel(parent)
ctx, cancel := context.WithCancel(parent)

// Signal untuk melakukan cancel
cancel()
```

## Context with deadline
```
var parent = context.Background()

ctx = context.WithCancel(parent)
ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))

// Signal untuk melakukan cancel
cancel()
```
