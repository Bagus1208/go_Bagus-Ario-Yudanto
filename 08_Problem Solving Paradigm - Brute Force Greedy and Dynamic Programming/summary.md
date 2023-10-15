# Summary

1. Problem solving paradigm adalah pendekatan yang biasanya digunakan untuk memecahkan masalah.

<br>

2. Ada 4 jenis problem solving paradigm yaitu :
   - Complete search(brute force)
   - Divide and Conquer
   - Greedy
   - Dynamic programming

<br>

3. Setiap jenis problem solving paradigm mempunyai keunggulan masing-masing. Untuk memecahkan suatu masalah kita perlu menentukan pendekatan yang paling sesuai dari ke 4 jenis problem solving problem tersebut.

# Problem Solving Paradigm
Problem solving paradigm adalah pendekatan yang biasa digunakan untuk menyelesaikan masalah. Setiap masalah perlu dipecahkan dengan pendekatan yang sesuai

## Complete search
* Complete search atau bruteforce metode untuk menyelesaikan masalah dengan mengecek seluruh ruang pencarian untuk mendapatkan solusi yang dibutuhkan
* Bruteforce digunakan jika tidak algoritma lain yang tersedia
* Biasanya mudah dibuat 
* Secara teori semua permasalahan dapat diselesaikan dengan pendekatan bruteforce khususnya ketika mempunyai waktu yang tak terbatas

Contoh penerapan complete search :
```
func linearSearch(elements []int,x int) int{
    for i:=0;i<len(elements);i++{
        if elements[i]==x{
            return i
        }
    }
    return -1
}
```
   
## Divide and conquer
Divide and conquer adalah paradigma penyelesaian masalah yang mana suatu permasalahan dibuat dengan cara yang lebih sederhana yaitu dengan cara membaginya menjadi bagian-bagian yang lebih kecil dan kemudian diselesaikan setiap bagian tersebut. 

Alur divide and conquer
1. Divide : membagi masalah yang besar menjadi masalah yang lebih kecil
2. Conquer : ketika masalah sudah cukup kecil untuk diselesaikan, maka langsung diselesaikan
3. Combine : jika dibutuhkan maka perlu menggabungkan solusi dari masalah-masalah yang lebih kecil menjadi solusi untuk masalah yang lebih besar

contoh penerapan divide and conquer :
```
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2

        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1 
}
```

## Greedy
Suatu algoritma dikatakan greedy jika algoritma tersebut membuat pilihan yang optimal secara lokal pada setiap langkahnya dengan harapan pada akhirnya mencapai solusi optimal secara global. dalam beberapa kasus, pekerjaan serakah, solusinya singkat dan berjalan efisien.

Contoh penerapan greedy :
* Huffman Coding
* Activity Selection
* Dijkstra algorithm - mencari jalur terpendek
   ```
   func Absolute(a, b int) int {
      if a < b {
         return b - a
      }

      return a - b
   }

   func Frog(jumps []int) int {
      var result int
      var rock int = 0

      for i := 0; i < len(jumps)/2; i++ {
         if rock == len(jumps)-2 {
            result += Absolute(jumps[rock+1], jumps[rock])
         } else if Absolute(jumps[rock+1], jumps[rock]) > Absolute(jumps[rock+2], jumps[rock]) {
            result += Absolute(jumps[rock+2], jumps[rock])
            rock += 2
         } else {
            result += Absolute(jumps[rock+1], jumps[rock])
            rock++
         }
      }

      return result
   }
   ```

## Dynamic programming
Dymanic programming adalah teknik algoritma yang digunakan untuk memecahkan masalah dengan cara memecahkannya menjadi submasalah yang lebih kecil dan menyimpan hasilnya untuk menghindari pengulangan perhitungan. Pendekatan ini digunakan untuk mengoptimalkan solusi masalah yang mempunyai sifat tumpang tindih (overlap) antara submasalah-submasalah yang dipecahkan.

Beberapa karakteristik utama dari Dynamic Programming adalah sebagai berikut:

* **Overlapping Subproblems (Submasalah yang Tumpang Tindih)**: Masalah besar sering kali dapat dipecahkan menjadi submasalah yang lebih kecil. Dynamic Programming bekerja dengan mengidentifikasi submasalah yang sama yang dihitung beberapa kali, dan menghitungnya hanya satu kali dan menyimpan hasilnya.

* **Optimal Substructure (Struktur Submasalah Optimal)**: Solusi optimal untuk masalah besar dapat dihasilkan dari solusi optimal dari submasalah yang lebih kecil.

* **Tabulation (Buat Tabel) atau Memoization (Penyimpanan Hasil)**: Terdapat dua pendekatan umum dalam Dynamic Programming. Pertama, tabulasi, yang melibatkan pembuatan tabel untuk menyimpan hasil perhitungan submasalah. Kedua, memoisasi, yang melibatkan penyimpanan hasil perhitungan submasalah dalam struktur data seperti memo atau cache.

contoh penerapan dynamic programming :
```
func fibonacci(number int) int {
	if number <= 1 {
		table[number] = number
		return number
	}

	if value, found := table[number]; found {
		return value
	}

	return fibonacci(number-1) + fibonacci(number-2)
}
```