
# Potentivio apps

![Logo](https://github.com/Hannon-App/be-api/blob/main/Hannon.png)

Readme in english [click here](https://github.com/Hannon-App/be-api/blob/main/README_english.md).

Hannon apps adalah sebuah aplikasi untuk peminjaman peralatan camping/ atribut peralatan outdoor yang menjembatani antara users (calon peminjam) dengan tenant (pihak yang meminjamkan). Dengan aplikasi ini, pihak tenant akan lebih mudah dan felxible dalam menjalankan bisnisnya untuk rent peralatan outdoornya, dan dari sisi users (calon peminjam) memudahkan untuk mencari items yang tersedia dari tenant dengan harga terjangkau karena tidak perlu untuk beli/ memiliki peralatan ketika akan melakukan kegiatan di outdoor (camping, hiking, dll). selain itu dengan adanya Hannon App kedepanya bisa menigkatkan intensitas pariwisata sehingga menjadi dampak baik untuk perkembangan pariwisata sehingga menghidupkan aktifitas ekonomi di wilayah tersebut.


## Fitur Users 

- Registers
- Login Users
- Edit Users
- melihat items & tenant
- Users mencari items dari tenant seusai dengan kota yang dituju
- Users dapat meminjam items dai tenant sesuai dengan waktu yang di tentukan oleh users.
- Users dapat membatalkan peminjaman yang sudah di boking dengan alasan yang valid
- Users dapat melakukan pemabayaran melalui xendit yang sudah di integrasikan
- Users dapat menerima notifikasi pembayaran melalui email aktif.

## Fitur Tenants

- Register
- Login Tenant
- Tenant dapat melakukan insert produk2nya yang akan dipinjamkan
- Tenant dapat mengedit/update data produk items nya
- Tenant dapat delete data produk items nya
- Tenant bisa melakukan transaksi dengan users

## Open APIs

Untuk Open API bisa lihat selengkapnya [disini](https://github.com/Hannon-App/be-api/blob/main/hannonapp-openAPI.yml)


## Menjalankan Lokal

Cloning project

```bash
  $ https://github.com/Hannon-App/be-api.git
```

Masuk ke direktori project

```bash
  $ cd ~/nama project kamu
```
Buat `database` baru

Buat sebuah file dengan nama di dalam folder root project `.env` dengan format dibawah ini. Sesuaikan configurasi di komputer lokal

```bash
export DBUSER='root'
export DBPASS='masukkan password kamu'
export DBHOST='127.0.0.1'
export DBPORT='3306'
export DBNAME='nama database kamu'
export JWTSECRET='......'
export KEY_API='......'
export KEY_API_SECRET='.........'
export CLOUD_NAME='.....'
export GOOGLE_APPLICATION_CREDENTIALS='keys.json'
export XENDIT_SECRET_KEY='massukan key secret dari xendit'
export CALLBACK_KEY='masuukan callback key dari xendit'
```

Jalankan aplikasi 

```bash
  $ go run main.go
```


## Authors

- [@firhanaf](https://github.com/firhanaf)
- [@royanqodri](https://github.com/royanqodri)
- [@Prayogarock](https://github.com/Prayogarock)

 
