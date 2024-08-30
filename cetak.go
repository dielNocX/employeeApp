package main

import "fmt"

func cetak(i int) {
	if exit == false {
		fmt.Println("[]=============================================================[]")
		fmt.Println("[][Data karyawan ke-", i+1, "]=======================================[]")
		fmt.Println("[]=============================================================[]")
		fmt.Printf("Nama					: %s\n", K[i].nama)
		fmt.Printf("NIK					: %s\n", K[i].NIK)
		fmt.Printf("Posisi					: %s\n", K[i].posisi)
		fmt.Printf("Departemen				: %s\n", K[i].departemen)
		fmt.Printf("Tanggal Masuk Karyawan			: %s\n", K[i].tglMasuk)
		fmt.Printf("Gender					: %s\n", K[i].gender)
		fmt.Println("[]=============================================================[]")
	}
}
