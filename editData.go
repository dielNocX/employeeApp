package main

import "fmt"

func editData(K *tabKaryawan, n int) {
	// IS: x adalah input yg dicari pada tabel dan y adalah opsi (1,2,3) untuk menentukan parameter yang dicari
	// FS:
	var indeks int = -1
	var opsi, nama_baru, NIK_baru, posisi_baru, departemen_baru, tglMasuk_baru, gender_baru, cari_NIK string

	if exit == false {
		fmt.Print("Masukkan data NIK karyawan yang ingin anda edit: ")
		fmt.Scan(&cari_NIK)
		indeks = cariIndeks(*K, n, 2, cari_NIK)
		if indeks >= 0 {
			cetak(indeks)
			fmt.Println("[]============================================[]")
			fmt.Println("[]             Edit Data Karyawan             []")
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
			fmt.Println("Data apa yang ingin anda edit?")
			fmt.Print("Masukkan inputan (0/6): ")
			fmt.Scan(&opsi)

			for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" && opsi != "6" && opsi != "0" {
				fmt.Println("Masukkan inputan yang valid pilih (0/6): ")
				fmt.Scan(&opsi)
			}
			if opsi == "1" {
				fmt.Println("Ganti Nama menjadi: ")
				fmt.Scan(&nama_baru)
				K[indeks].nama = nama_baru
			} else if opsi == "2" {
				fmt.Println("Ganti NIK menjadi: ")
				fmt.Scan(&NIK_baru)
				if cek(*K, NIK_baru) == false {
					fmt.Println("NIK telah terpakai, masukkan NIK yang berbeda")
					fmt.Print("\nGanti NIK menjadi: ")
					fmt.Scan(&NIK_baru)
				}
				K[indeks].NIK = NIK_baru
			} else if opsi == "3" {
				fmt.Println("Ganti Posisi menjadi: ")
				fmt.Scan(&posisi_baru)
				K[indeks].posisi = posisi_baru
			} else if opsi == "4" {
				fmt.Println("Ganti Departemen menjadi: ")
				fmt.Scan(&departemen_baru)
				K[indeks].departemen = departemen_baru
			} else if opsi == "5" {
				fmt.Println("Ganti Tanggal Masuk Karyawan Menjadi: ")
				fmt.Scan(&tglMasuk_baru)
				K[indeks].tglMasuk = tglMasuk_baru
			} else if opsi == "6" {
				fmt.Println("Ganti Gender menjadi: ")
				fmt.Scan(&gender_baru)
				K[indeks].gender = gender_baru
			} else if opsi == "0" {
				menu()
			}
			cetak(indeks)
		} else {
			dataNotFound()
		}
		menu()
	}
}
