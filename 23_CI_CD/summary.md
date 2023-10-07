# CI/CD
## Continuous Integration (CI)
Continuous Integration (CI) adalah praktik dalam pengembangan perangkat lunak yang mengacu pada penggabungan, pengujian, dan pembangunan kode secara otomatis dan berulang. Tujuannya adalah untuk meningkatkan kualitas perangkat lunak dengan memastikan bahwa perubahan kode yang baru diintegrasikan ke dalam repositori utama (main branch) diuji secara otomatis dan segera. 

Siklus Continuous Integration :
* **Integrasi Kode**: Pengembang mengintegrasikan kode baru yang mereka tulis ke dalam basis kode setidaknya sekali sehari. Ini dilakukan untuk menjaga salinan lokal Anda agar tidak terlalu jauh dari cabang utama dari pembuatan kode.

* **Pemeriksaan Otomatis**: Setelah kode baru ditambahkan, pengujian otomatis dilakukan terhadap setiap iterasi dari build untuk mengidentifikasi masalah integrasi lebih awal, ketika mereka lebih mudah untuk diperbaiki1. Ini membantu Anda menghindari konflik gabungan yang bisa “merusak” build dan membutuhkan tim berjam-jam atau berhari-hari untuk menyelesaikannya.

* **Pemberitahuan**: Seluruh tim pengembangan akan tahu dalam hitungan menit setelah check-in apakah Anda telah membuat kode yang buruk, karena layanan integrasi terus menerus secara otomatis membangun dan menguji perubahan kode Anda untuk setiap kesalahan.

## Continuous Deployment/Delivery (CD)
Continuous deployment adalah pendekatan pengembangan perangkat lunak di mana perubahan kode pada aplikasi diterapkan secara otomatis ke lingkungan produksi. Hal ini dilakukan dengan menggunakan serangkaian tes yang telah ditentukan sebelumnya. Tujuan dari continuous deployment adalah untuk memfasilitasi rilis yang lebih cepat dengan menggunakan otomatisasi untuk mengurangi kebutuhan intervensi manusia sebanyak mungkin selama proses penyebaran. 

Siklus Continuous Deployment/Delivery :
* **Unit Test**: Di sini, kode yang ditulis oleh pengembang diuji pada tingkat unit untuk memastikan bahwa setiap bagian kode berfungsi seperti yang diharapkan.

* **Platform Test**: Setelah tes unit, kode diuji pada platform tertentu untuk memastikan bahwa kode berfungsi dengan baik pada platform tersebut.

* **Deliver to Acceptance**: Setelah kode lolos tes platform, kode dikirim ke lingkungan penerimaan di mana kode tersebut diuji lebih lanjut.

* **Application Production**: Jika kode lolos tes penerimaan, maka kode tersebut siap untuk diproduksi.

* **Deploy to Production**: Pada tahap ini, kode yang telah diproduksi diterapkan ke lingkungan produksi.

* **Post Deploy Tests**: Setelah penyebaran ke produksi, tes pasca penyebaran dilakukan untuk memastikan bahwa semua fitur berfungsi dengan baik di lingkungan produksi.

## Tools untuk CI/CD
* **Jenkins**: Ini adalah alat CI/CD berbasis Java yang bersifat open-source. Jenkins menawarkan integrasi terus-menerus, bersama dengan memfasilitasi pengiriman terus-menerus. Jenkins juga memungkinkan pengujian dan pelaporan real-time.

* **CircleCI**: CircleCI adalah alat CI/CD yang mendukung pengembangan dan penerbitan perangkat lunak yang cepat. CircleCI memungkinkan otomatisasi di seluruh pipeline pengguna, mulai dari pembuatan kode, pengujian hingga penyebaran.

* **AWS CodeBuild**: Ini adalah layanan integrasi terus-menerus yang mengkompilasi kode sumber, menjalankan tes, dan menghasilkan paket perangkat lunak yang siap untuk dikerahkan.

* **Azure DevOps**: Ini adalah layanan Microsoft yang menyediakan alat pengembangan untuk merencanakan pekerjaan, berkolaborasi pada kode, dan menerapkan aplikasi.

* **TeamCity**: Ini adalah server integrasi terus-menerus yang dirancang oleh JetBrains. TeamCity mendukung berbagai alat dan kerangka kerja, termasuk tetapi tidak terbatas pada .NET, Java, Ruby, Python, PHP, dan banyak lagi.

* **Bamboo**: Ini adalah server CI/CD dari Atlassian yang memungkinkan otomatisasi build, tes, dan rilis dalam satu tempat.