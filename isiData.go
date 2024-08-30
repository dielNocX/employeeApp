package main

import "fmt"

func isiData(K *tabKaryawan, n *int) {
	// IS: Array A dengan banyak elemen n terdefinisi sembarang
	// Proses: Membaca bilangan bulat dan memasukkan bilangan bulat positif ke dalam array A. Pembacaan dilakukan selama terbaca bilangan positif dan n < NMAX.
	// FS: Array A sebanyak n elemen berisi nilai
	var nama, NIK, posisi, departemen, tglMasuk, gender, konfirmasi string

	
	if exit == false {
		fmt.Println("[]============================================[]")
		fmt.Print("Nama			: ")
		fmt.Scan(&nama)

		if nama != "1" {
			fmt.Print("NIK			: ")
			fmt.Scan(&NIK)
			fmt.Print("Posisi			: ")
			fmt.Scan(&posisi)
			fmt.Print("Departemen		: ")
			fmt.Scan(&departemen)
			fmt.Print("Tanggal Masuk Karyawan	: ")
			fmt.Scan(&tglMasuk)
			fmt.Print("Gender			: ")
			fmt.Scan(&gender)
			for cek(*K, NIK) == false {
				fmt.Println("NIK telah terpakai, masukkan NIK yang berbeda")
				fmt.Print("NIK			: ")
				fmt.Scan(&NIK)
			}
		}
		for nama != "." && *n < NMAX {

			K[*n] = karyawan{
				nama:       nama,
				NIK:        NIK,
				posisi:     posisi,
				departemen: departemen,
				tglMasuk:   tglMasuk,
				gender:     gender,
				cuti:       cutiKaryawan{saldo: 12}}
			if *n < NMAX {
				*n++

			}
			fmt.Println("[]================[KONFIRMASI]================[]")
			fmt.Println("[]====[]======================================[]")
			fmt.Println("[] 1. [] Lanjutkan isi data karyawan          []")
			fmt.Println("[]====[]======================================[]")
			fmt.Println("[] 0. [] Kembali                              []")
			fmt.Println("[]====[]======================================[]")
			fmt.Print("Masukkan inputan (0/1): ")
			fmt.Scan(&konfirmasi)
			for konfirmasi != "1" && konfirmasi != "0" {
				fmt.Println("Masukkan inputan yang valid (0/1)")
				fmt.Scan(&konfirmasi)
			}
			if konfirmasi != "0" && konfirmasi == "1" {
				fmt.Println("[]============================================[]")
				fmt.Print("Nama			: ")
				fmt.Scan(&nama)
				fmt.Print("NIK			: ")
				fmt.Scan(&NIK)
				fmt.Print("Posisi			: ")
				fmt.Scan(&posisi)
				fmt.Print("Departemen		: ")
				fmt.Scan(&departemen)
				fmt.Print("Tanggal Masuk Karyawan	: ")
				fmt.Scan(&tglMasuk)
				fmt.Print("Gender			: ")
				fmt.Scan(&gender)
				fmt.Println("[]============================================[]")
				cariIndeks(*K, *n, 2, NIK)
				for cek(*K, NIK) == false {
					fmt.Println("NIK telah terpakai, masukkan NIK yang berbeda")
					fmt.Print("NIK			: ")
					fmt.Scan(&NIK)
				}
			} else {
				menuIsiData()
			}

		}
	}
}
