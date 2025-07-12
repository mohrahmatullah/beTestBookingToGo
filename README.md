# beTestBookingToGo

Project ini adalah backend service untuk BookingToGo Test.

## Persyaratan

Sebelum menjalankan aplikasi, pastikan Anda sudah menginstall:

- **Golang** versi **1.21.4**

- **PostgreSQL** versi **14.5** atau lebih baru

## Konfigurasi Database

1. Pastikan PostgreSQL sudah berjalan di mesin Anda.
2. Buat database dengan nama `beTestBookingToGo`.
3. Edit file `config.json` sesuai pengaturan koneksi database Anda:

```json
	{
		"connection_string": "host=127.0.0.1 port=5432 user=postgres dbname=beTestBookingToGo sslmode=disable",
		"port": 8081
	}
```

```json
host: Alamat server PostgreSQL
port: Port PostgreSQL (default 5432)
user: Username PostgreSQL
dbname: Nama database
sslmode: Mode SSL (disable jika lokal)
```

## Cara Menjalankan Aplikasi

``json
go run ./cmd 
```