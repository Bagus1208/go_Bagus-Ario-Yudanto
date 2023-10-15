# Summary

1. Join berfungsi untuk menampilkan data dari beberapa tabel yang memiliki relasi sekaligus. Namun jika memakainya terlalu banyak dapat membuat performa semakin menurun. Join ada beberapa jenis yaitu :
   * Inner join
   * Left join
   * Right join

<br>

2. Subquery adalah query yang berada didalam query lain. Fungsi subquery adalah untuk mengembalikan data yang digunakan untuk batasan data yang akan diambil dari query utama. Contoh subqury : 
    * SELECT * FROM products WHERE harga > (SELECT MIN(harga) FROM products);

<br>

3. Function bisa mengembalikan sebuah nilai pada pemanggilnya dan juga bisa menjalankan query dengan kondisi tertentu yang biasa disebut trigger function.

# Join, Unioin, Agregasi, Subquery, Function
## JOIN
JOIN digunakan untuk menggabungkan data dari dua atau lebih tabel berdasarkan kolom-kolom yang memiliki nilai yang sama.

### INNER JOIN 
* Inner join menggabungkan baris dari kedua tabel berdasarkan nilai yang cocok di kolom yang ditentukan.
* Hanya baris yang memiliki nilai yang cocok dalam kedua tabel akan disertakan dalam hasil.
  
Contoh penggunaan Inner Join:
```
SELECT orders.order_id, customers.customer_name
FROM orders
INNER JOIN customers ON orders.customer_id = customers.customer_id;
```

### LEFT JOIN
* Left join menggabungkan semua baris dari tabel kiri (pertama) dengan baris yang cocok dari tabel kanan (kedua).
* Jika tidak ada cocokan dalam tabel kanan, kolom-kolom di tabel kanan akan memiliki nilai NULL dalam hasil.

Contoh penggunaan Left Join:
```
SELECT employees.employee_name, departments.department_name
FROM employees
LEFT JOIN departments ON employees.department_id = departments.department_id;
```

### Right Join
* Right join adalah kebalikan dari left join. Ini menggabungkan semua baris dari tabel kanan dengan baris yang cocok dari tabel kiri.
* Jika tidak ada cocokan dalam tabel kiri, kolom-kolom di tabel kiri akan memiliki nilai NULL dalam hasil.
  
Contoh penggunaan Right Join:
```
SELECT employees.employee_name, departments.department_name
FROM employees
RIGHT JOIN departments ON employees.department_id = departments.department_id;
```

## Union
UNION adalah operasi dalam SQL yang digunakan untuk menggabungkan hasil dari dua atau lebih query menjadi satu hasil tunggal. Ini menggabungkan semua baris unik dari setiap query dan mengembalikan satu set data yang tidak berisi duplikat.

Contoh penggunaan union:
```
SELECT City FROM customers
UNION
SELECT City FROM suppliers;
```

## Fungsi Agregat
Fungsi agregat dalam database digunakan untuk melakukan perhitungan statistik atau operasi matematika pada kumpulan data. Fungsi-fungsi agregat ini memungkinkan Anda untuk mengambil informasi ringkasan dari data dalam tabel.

Contoh fungsi agregat COOUNT():
```
SELECT COUNT(*) FROM employees;
```

## Subquery
Subquery, juga dikenal sebagai nested query atau inner query, adalah query yang ditempatkan di dalam query utama. Subquery digunakan untuk mengambil data yang diperlukan untuk kondisi dalam query utama. Subquery dapat digunakan di dalam pernyataan SELECT, FROM, WHERE, dan banyak lagi.

Contoh subquery: 
```
SELECT product_name
FROM products
WHERE category_id = (SELECT category_id FROM categories WHERE category_name = 'Electronics');
```
## Function
Function adalah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya

Contoh function:
```
DELIMITER $$
CREATE PROCEDURE GetEmployeeName(IN employee_id INT)
BEGIN
  SELECT first_name, last_name
  FROM employees
  WHERE employee_id = employee_id;
END$$
DELIMITER ;
```