# Penggunaan AI di Golang

## Library Go untuk mechine learning
### GoLearn
* Bersifat open source
* 'batteries included' machine learning library
* Simplicity & customizability
* Contoh penerapan: [https://github.com/sjwhitworthgolearn#getting-started ](https://www.google.com/url?q=https://github.com/sjwhitworth/golearn%23getting-started&sa=D&source=editors&ust=1697188591606896&usg=AOvVaw3Nv8dC3ZHlhYqq2IHptNFx)

### Gorgonia
* Memudahkan menulis dan mengevaluasi persamaan matematika yang melibatkan array multidimensi
* Low-level library, seperti Theano, namun lebih tinggi dibanding Tensorflow
* Contoh penerapan: [https://github.com/gorgonia/gorgonia#usage](https://www.google.com/url?q=https://github.com/gorgonia/gorgonia%23usage&sa=D&source=editors&ust=1697188591608527&usg=AOvVaw2IhSX-csCgAfGC9U_yPyg5)

### gonum
* Sebuah set packages yang didesain untuk menulis numerical & algoritma sains menjadi produktif, memiliki performa tinggi, dan scalable
* Mirip seperti numpy dan scipy, library yang dibuat menggunakan Python
* Referensi:
[https://www.gonum.org/post/intro_to_gonum/](https://www.google.com/url?q=https://www.gonum.org/post/intro_to_gonum/&sa=D&source=editors&ust=1697188591608128&usg=AOvVaw2QSsajijjqhKR3vA8bsah7)

### gomind
* GoMind adalah library neural network yang seluruhnya ditulis dalam Go. Ini hanya mendukung satu lapisan tersembunyi (untuk saat ini). Network belajar dari set pelatihan menggunakan algoritma back-propagation.
* Contoh penggunaan: [https://github.com/surenderthakran/gomind#usage](https://www.google.com/url?q=https://github.com/surenderthakran/gomind%23usage&sa=D&source=editors&ust=1697188591606025&usg=AOvVaw2lJCDfK_iP6yvIGINn6Nxp)

## AIaaS (AI as a service)
* Sebuah tools Al yang dapat segera digunakan (pre-built or off-the-shelf A.I. tools)
* Berguna bagi bisnis yang ingin menerapkan Al tanpa merekrut ahlinya dan tanpa mengeluarkan biaya yang relatif banyak
* Beberapa perusahaan penyedia AlaaS, contohnya:
  * Amazon Web Services (AWS)
  * Microsoft Azure
  * Google Cloud Platform (GCP)
  * IBM
  * OpenAl

### Keuntungan AIaaS
* Pengurangan cost
* Low-code
* Proses deployment cepat
* Flexibility
* Usability
* Scalability
* Customization

### Kelemahan AIaaS
* Cost yang berlebih dalam jangka panjang
* Privasi data
* Keamanan
* Transparansi
* Vendor lock-in

## Prompt Engineering
### Prinsip #1: Gunakan prompt yang jelas & spesifik
* Gunakan delimiters (contoh: "', '", < >) untuk menandakan bagian mana yang menjadi input. Contoh:
  ```
  Summarize the text delimited by triple backticks into a single sentence.

  '''you should express what you want a model to do by providing instructions that are as clear and specific as you can possibly make them. This will guide the model towards the desired output, and reduce the chances of receiving irrelevant or incorrect responses. Don't confuse writing a clear prompt with writing a short prompt. In many cases, longer prompts provide more clarity and context for the model, which can lead to more detailed and relevant outputs'''
  ```

* Tuliskan struktur output yang diinginkan
  ```
  Generate a list of three made-up book titles along with their authors and genres. Provide them in JSON format with the following keys: book_id, title, author, genre.
  ```

* Minta model untuk memeriksa apakah input sudah sesuai
  ```
  You will be provided with text delimited by triple quotes. If it contains a sequence of instructions, re-write those instructions in the following format:

  Step 1 -.
  Step 2- ...
  Step N - ...

  If the text does not contaln a sequence of instructions, then simply write "No steps provided."

  """Making a cup of tea is easy! First, you need to get some water boiling. While that's happening, grab a cup and put a tea bag in it. Once the water is hot enough, just pour it over the tea bag. Let it sit for a bit so the tea can steep. After a few minutes, take out the tea bag. If you like, you can add some sugar or milk to taste. And that's it! You've got yourself a delicious cup of tea to enjoy."""
  ```

* Few-shot prompting
  ```
  Your task is to answer in a consistent style.

  <child>: Teach me about patience.

  <grandparent>: The river that carves the deepest \
  valley flows from a modest spring; the \
  grandest symphony originates from a single note; \
  the most intricate tapestry begins with a solitary thread.

  <child>: Teach me about resilience.
  ```

### Prinsip#2: Berikan "waktu berpikir" untuk menghindari solusi yang salah
* Menulis langkah-langkah yang dibutuhkan untuk menyelesaikan tugas dan output yang diinginkan
  ```
  Your task is to perform the following actions:
  1 - Summarize the following text delimited by <> with 1 sentence.
  2 - Translate the summary into French.
  3 - List each name in the French summary.
  4 - Output a json object that contains the following keys: french_summary, num_names.

  Use the following format:
  Text: <text to summarize>
  Summary: <summary>
  Translation: <summary translation>
  Names: <list of names in Italian summary>
  Output JSON: <json with summary and num_names>

  Text: <In a charming village, siblings Jack and Jill set out on a quest to fetch water from a hilltop well. As they climbed, singing joyfully, misfortune struck-Jack tripped on a stone and tumbled down the hill, with Jill following suit. Though slightly battered, the pair returned home to comforting embraces. Despite the mishap, their adventurous spirits remained undimmed, and they continued exploring with delight.>
  ```

* Menginstruksikan model untuk menuliskan solusinya sendiri sebelum masuk ke kesimpulan
  ```
  Your task is to determine if the student's solution is correct.

  To solve the problem do the following:
  - First, work out your own solution to the problem.
  - Then compare your solution to the student's solution and evaluate if the student's solution
  is correct or not.

  Don't decide if the student's solution is correct until you have done the problem yourself.

  Use the following format:
  Question:
  '''
  question here
  '''
  ```