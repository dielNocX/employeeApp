package main

import "fmt"

func deleteData(K *tabKaryawan, n *int) {
	var cari_NIK, opsi string
	var delete karyawan
	var indeks int = -1
	if exit == false {
		fmt.Println("Masukkan NIK yang ingin anda delete: ")
		fmt.Scan(&cari_NIK)
		indeks = cariIndeks(*K, *n, 2, cari_NIK)
		if indeks >= 0 {
			cetak(indeks)
			fmt.Println("Apakah anda yakin ingin delete data karyawan berikut? (ya/tidak)")
			fmt.Scan(&opsi)
			for opsi != "ya" && opsi != "Ya" && opsi != "yA" {
				fmt.Println("Masukkan inputan yang valid (Ya/Tidak): ")
				fmt.Scan(&opsi)
			}
			if opsi == "ya" || opsi == "Ya" || opsi == "yA" {
				for indeks < *n-1 {
					K[indeks] = K[indeks+1]
					indeks++
				}
				K[indeks] = delete
				*n -= 1
				fmt.Println("[Data berhasil di hapus]")
			}
		} else {
			dataNotFound()
		}
		menu()
	}
}
