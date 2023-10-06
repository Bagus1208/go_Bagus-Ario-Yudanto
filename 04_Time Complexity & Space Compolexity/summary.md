# Summary

1. Kompleksitas program dibagi menjadi 2 yaitu :
   1. Time complexity adalah seberapa lama waktu yang dibutuhkan untuk menjalankan suatu algoritma. 
   2. Space complexity adalah seberapa besar memori yang digunakan dalam menjalankan suatu algoritma.
   
<br>

2. Big O notation adalah cara untuk mengkonversi dominant operation kedalam bentuk aljabar. Dominant operation adalah operasi primitif (pertambahan, perkalian, dll) yang berdampak besar terhadap kompleksitas algoritma.

<br>

3. Macam-macam time complexity
   * Constant time - O(1)
      * Berapapun input yang dimasukan dominant operation akan dijalankan 1 kali
   * Linier time - O(n)
      * Dominant operation akan dijalankan sebanyak n kali
   * Logarithmic time - O(log n)
      * Jika kita menginput sebesar n terhadap sebuah fungsi, dominant operation yang dijalankan sebuah fungsi akan berkurang berdasarkan suatu faktor 
   * Quadratic time - O(n^2)
      * Input ukuran n, dominant operation yang diperlukan untuk menyelesaikan tugas adalah kuadrat dari n
      * Input ukuran n, dominant operation yang diperlukan untuk menyelesaikan tugas adalah konstanta pangkat n

# Time Complexity & Space Complexity
## Time Complexity
Time complexity digunakan untuk mempermudah dalam menghitung estimasi waktu eksekusi program. Kompleksitas dalam dilihat sebagai jumlah maksimal operasi primitif yang dijalankan program. contoh operasi primitif yaitu pertambahan, perkalian, penugasan, dll. Beberapa operasi mungkin tidak dihitung dan fokus diberikan pada operasi yang dilakukan sebanyak mungkin. Operasi seperti ini disebut Dominan.

Contoh operasi dominan :
```
func dominant(n int) int {
   var result int = 0
   for i := 0; i < n; i++ {
      result += i   // <-- ini adalah operasi dominan
   }
   Result += result + 10
   return result
}
```

### Jenis-jenis time complexity
1. Constant time - O(1) 
   ```
   func dominant(n int) int {
      var result int = 0 
      result = result + 10
      return result 
   }
   ```
2. Linier time - O(n)
   ```
   func linier(n int, A[]int) int {
      for i := 0; i < n; i++ {
         if A[i] == 0 {
            return 0
         }
      }
      return 1
   }
3. Linier time - O(m + n)
   ```
   func linier2(n int, m int) int {
      var result = 0
      fun i := 0; i < n; i++ {
         result += 1
      }
      func j := 0; j < m; j++ {
         result += 1
      }
      return result
   }
   ```
4. Logarithmic time - O(log n)
   ```
   func logarithmic(n int) int {
      var result = 0
      for n > 1 {
         n /= 2
         result += 1
      }
      return result
   }
5. Quadratic time - O(n^2)
   ```
   func quadratic(n int) int {
      for i := 0; i < n; i++ {
         for j := i; j < n; j++ {
            result += 1
         }
      }
      return result
   }

## Time Limit
Saat ini, rata-rata komputer dapat menjalankan 10^8 operasi dibawah satu detik. Time limit yang ditetapkan dalam tes online biasanya dari 1 sampai 10 detik.
* n <= 1.000.000, time complexity yang diharapkan yaitu O(n) or O(n log n)
* n <= 10.000, time complexity yang diharapkan yaitu O(n^2)
* n <= 500, time complexity yang diharapkan yaitu O(n^3)

## Space Complexity
Batasan memori memberikan informasi tentang space complexity yang diharapkan. Anda dapat mengestimasi jumlah variabel yang dapat kamu buat dalam program. Secara singkat, jika Anda memiliki jumlah variabel yang tetap, Anda juga memiliki space complexity yang tetap. Dalam notasi Big-O, ini adalah O(1). Jika Anda perlu mendeklarasikan sebuah array dengan n elemen, Anda memiliki space complexity linier - O(n).