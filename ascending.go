package main

import "fmt"

func ascending(K *tabKaryawan, n int) {
	var opsi string
	if exit == false {
		fmt.Println("[]============================================[]")
		fmt.Println("[]     Urutkan Data Karyawan Berdasarkan      []")
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
			ascendingSort(K, n, 1)
		} else if opsi == "2" {
			ascendingSort(K, n, 2)
		} else if opsi == "3" {
			ascendingSort(K, n, 3)
		} else if opsi == "4" {
			ascendingSort(K, n, 4)
		} else if opsi == "5" {
			ascendingSort(K, n, 5)
		} else if opsi == "6" {
			ascendingSort(K, n, 6)
		} else if opsi == "0" {
			menuMenampilkanData()
		}
		menampilkanData(*K, n)
	}
}
