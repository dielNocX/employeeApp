package main

import "fmt"

func menampilkanData(k tabKaryawan, n int) {
	var konfirmasi string
	if exit == false {
		for i := 0; i < n; i++ {
			fmt.Println("[]=============================================================[]")
			fmt.Println("[][Data karyawan ke-", i+1, "]=======================================[]")
			fmt.Println("[]=============================================================[]")
			fmt.Printf("Nama					: %s\n", k[i].nama)
			fmt.Printf("NIK					: %s\n", k[i].NIK)
			fmt.Printf("Posisi					: %s\n", k[i].posisi)
			fmt.Printf("Departemen				: %s\n", k[i].departemen)
			fmt.Printf("Tanggal Masuk Karyawan			: %s\n", k[i].tglMasuk)
			fmt.Printf("Gender					: %s\n", k[i].gender)
			if i == n-1 {
				fmt.Println("[]=============================================================[]")
			}
		}
		if k[0].nama == "" {
			dataNotFound()
			menu()
		}
		fmt.Println()
		fmt.Println("[]================[KONFIRMASI]================[]")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 1. [] Urutkan data karyawan                []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 0. [] Kembali                              []")
		fmt.Println("[]====[]======================================[]")
		fmt.Print("Masukkan inputan (0/1): ")
		fmt.Scan(&konfirmasi)
		for konfirmasi != "1" && konfirmasi != "0" {
			fmt.Println("Masukkan inputan yang valid (0/1)")
			fmt.Scan(&konfirmasi)
		}
		if konfirmasi == "1" {
			menuDataTerurut()
		} else if konfirmasi == "0" {
			menuMenampilkanData()
		}
	}
}
