# Summary

1. Rekursif adalah fungsi yang di dalam bagiannya memanggil dirinya sendiri berulang kali sampai keadaan tertentu. Didalam rekursif ada dua bagian yaitu :
   1. Base case 
      - Base case adalah bagian yang berisi nilai dasar dari proses rekursif dan menjadi kondisi berhentinya proses rekursi
   2. Recursive case
      - Recursive case adalah bagian fungsi yang memanggil dirinya sendiri 

<br>

2. Serching adalah proses untuk mencari satu nilai dalam suatu kumpulan nilai. Contoh searching yaitu :
   1. Linier search
   2. binary search

<br>

3. Sorting adalah proses untuk mengurutkan suatu kumpulan data dengan urutan tertentu. Contoh sorting yaitu :
   1. Bubble sort
   2. Selection sort
   3. Insert sort
   4. Merge sort 
   5. Quick sort
   6. Counting sort

# Recursive, Number Theory, Sorting, Searching
## Recursive
Recursive adalah konsep dalam pemrograman di mana sebuah fungsi dapat memanggil dirinya sendiri selama proses eksekusi. Dalam rekursi, sebuah masalah besar dipecahkan menjadi masalah-masalah yang lebih kecil dengan cara yang serupa. Fungsi rekursif akan terus memanggil dirinya sendiri sampai mencapai kondisi terminasi atau base case, di mana rekursi berhenti dan nilai-nilai yang dikumpulkan dikembalikan.

### Alasan menggunakan recursive
* Beberapa masalah dapat dengan mudah diselesaikan dan kode lebih pendek dengan menggunakan pendekatan recursive
* Pada dasarnya recursive dan looping melakukan perulangan
* Namun, terkadang untuk memecahkan masalah menggunakan solusi perulangan sulit untuk dipikirkan dan membutuh kan teknik khusus
* Dengan menggunakan recursive, mungkin lebih mudah untuk dilihat dan mendesain jalur penyelesaian

### Strategi recursive
Ada dua hal yang perlu dipikirkan ketika menggunakan strategi recursive
* Base case : apa kasus sederhana dari permasalahan ini
* Recurrance relations : apa hubungan recursive dari masalah ini dengan masalah serupa yang lebih kecil

Contoh program recursive
```
func factorial(number string) {
   if number == 1 {
      return 1
   }

   return n * factorial(number - 1)
}
```

## Number of Theory
Number of theory adalah cabang matematika yang mempelajari bilangan bulat. Banyak sekali topik dalam number of theory seperti bilangan prima, faktor persekutuan terbesar, kelipatan persekutuan terkecil, faktorial, faktor prima, dll

### Faktorial 
```
func factorial(number string) {
   if number == 1 {
      return 1
   }

   return n * factorial(number - 1)
}
```

### Bilangan prima
```
func checkPrime(number) int {
   if number < 2 {
      return false
   } 

   for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
```

### Faktor persekutuan terbesar
```
func gcd(a, b int) int {
   if b == 0 {
      return a
   }

   return gcd(b, a % b)
}
```

### Kelipatan persekukutan terkecil
```
func lcm(a, b int) int {
   return a * (b / gcd(a, b))
}
```

## Searching
Searching adalah proses mencari nilai tertentu dalam sebuah kelompok nilai

### Linier search - O(n)
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

### Builtins search
```
elements := []int{12,15,15,19,24,31,53,59,60}
value := 31
findIndex := sort.SearchInts(elements, value)
if value == elemnts[findIndex]{
    fmt.println("Value found in elements")
}else{
    fmt.println("Value not found in elements")
}
```

## Sorting
Sorting adalah proses mengurutkan data dengan urutan tertentu, biasanya mengurutkan berdasarkan nilai elemen. Kita dapat mengurutkan angka, kata, pasangan, dan masih banyak lagi. Misalnya kita dapat mengurutkan siswa berdasarkan tinggi badannya, dan kita dapat mengurutkan kota menurut abjad atau menurut jumlahwarganya, urutan yang paling sering digunakan adalah urutan angka dan urutan abjad

### Selection sort - O(n^2)
```
func selectionSort(arr []int) {
    n := len(arr)

    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        arr[i], arr[minIndex] = arr[minIndex], arr[i] // Swap the minimum element with the current element
    }
}
```

### Counting sort - O(n + k)
```
func countingSort(arr []int) []int {
	count := make([]int, k+1)
   for i := 0; i < len(elements); i++ {
      count[elements[i]]++
   }
   counter := 0
   for i := 0; i < k; i ++ {
      for j := 0; j < count[i]; j++ {
         elements[counter] = i
         counter += 1
      }
   }

   return elements
}
```   

### Builtins sort
```
strs := []string{"x", "g", "u"}
sort.Strings(strs)
fmt.Println(strs)       // [g, u, x]

ints := []int{4, 7, 1}
sort.Ints(ints)
fmt.Println(ints)       // [1, 4, 7]

s := sort.IntsAreSorted(ints)
fmt.Println(s)          // true