 Golang Simple Boilerplate with Gofiber & Gorm
## Boilerplate Structure
- config/
- controller/
- entity/
- exception/
- helper/
- middleware/
- model/
- repository/
- seed/
- service/
- test/
- validation/
- .env.example
- .env
- .env.test
- .gitignore
- main.go
- README.md

## Definisi
- config/ - Konfigurasi API, contohnya konfigurasi koneksi Elasticsearch, Mysql, MongoDB, Redis, Kafka, etc.
- controller/ - Tempat Controller semua route Gofiber, contohnya mengatur endpoint, request ke Service, etc.
- entity/ - Entity yang biasanya berisi field-field yang akan digunakan/dipanggil dari database.
- contohnya  insert/get data dari redis, ubah type string ke type int, etc.
- middleware/ - Mirip seperti helper, middleware biasa digunakan pada Controller untuk diproses sebelum controller berjalan/selesai.
- model/ - Mirip seperti entity, model biasanya digunakan untuk memberikan response atau sebagai value return.
- repository/ - Tempat proses transaksi dari/ke database.
- seed/ - 
- service/ - Sebagai jembatan antara Controller dan Service. Biasa digunakan untuk handle/validasi data sebelum dilanjut ke Repository.
- test/ - Tempat file-file Unit Test Golang.
- validation/ - Mirip seperti helper, validation biasanya digunakan/dipanggil di Service.
- .env.example - Contoh/format environment untuk disalin
- .env - Environment utama/production
- .env.test - Environment untuk testing
- .gitignore - Ignore file/folder saat push ke git
- main.go - File utama/main Golang, tempat panggil Repository, Service, Controller.
- README.md - File README

## Trivial & Tips
- Lebih baik gunakan Service untuk logic sebelum masuk ke repository.
- Lebih baik gunakan Struct dari Model untuk Response, supaya lebih mudah penggunaannya.
- Selalu gunakan .env.test untuk Unit Testing.


