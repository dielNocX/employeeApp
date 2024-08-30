package main

import "fmt"

func pilihKaryawan(K tabKaryawan, n int, idx *int) {

	// var opsi, nama_baru, NIK_baru, posisi_baru, departemen_baru, tglMasuk_baru, gender_baru, cari_NIK string
	var cari_NIK string
	if exit == false {
		*idx = -1
		fmt.Print("Masukkan NIK karyawan : ")
		fmt.Scan(&cari_NIK)

		*idx = cariIndeks(K, n, 2, cari_NIK)
		if *idx >= 0 {
			cetak(*idx)
		} else {
			dataNotFound()
		}
		// fmt.Print("Masukkan data NIK karyawan yang ingin anda edit: ")
	}
}
