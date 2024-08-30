package main

import "fmt"

func cariData(K tabKaryawan, n int) {
	var opsi, x string
	var indeks int = -1
	if exit == false {
		fmt.Println("[]============================================[]")
		fmt.Println("[]       Cari Data Karyawan Berdasarkan       []")
		fmt.Println("[]============================================[]")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 1. [] Nama                                 []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 2. [] NIK                                  []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 3. [] Posisi                               []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 4. [] Departemen                           []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 5. [] Tanggal Masuk Karyawan               []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 6. [] Gender                               []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 0. [] Kembali                              []")
		fmt.Println("[]====[]======================================[]")
		fmt.Print("Masukkan inputan (0/6): ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" && opsi != "6" && opsi != "0" {
			fmt.Println("Masukkan inputan yang valid (0/6)")
			fmt.Scan(&opsi)
		}
		if opsi == "1" {
			fmt.Println("Masukkan data nama karyawan yang ingin anda cari: ")
			fmt.Scan(&x)
			indeks = cariIndeks(K, n, 1, x)
		} else if opsi == "2" {
			fmt.Println("Masukkan data NIK karyawan yang ingin anda cari: ")
			fmt.Scan(&x)
			indeks = cariIndeks(K, n, 2, x)
		} else if opsi == "3" {
			fmt.Println("Masukkan data Posisi karyawan yang ingin anda cari: ")
			fmt.Scan(&x)
			indeks = cariIndeks(K, n, 3, x)
		} else if opsi == "4" {
			fmt.Println("Masukkan data Departemen karyawan yang ingin anda cari: ")
			fmt.Scan(&x)
			indeks = cariIndeks(K, n, 4, x)
		} else if opsi == "5" {
			fmt.Println("Masukkan data Tanggal Masuk Karyawan yang ingin anda cari: ")
			fmt.Scan(&x)
			indeks = cariIndeks(K, n, 5, x)
		} else if opsi == "6" {
			fmt.Println("Masukkan data Gender karyawan yang ingin anda cari: ")
			fmt.Scan(&x)
			indeks = cariIndeks(K, n, 6, x)
		} else if opsi == "0" {
			menu()
		}
		if indeks >= 0 {
			cetak(indeks)
		} else {
			dataNotFound()
		}
		menu()
	}
}
