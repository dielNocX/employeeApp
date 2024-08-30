package main

import "fmt"

func menuIsiData() {
	var opsi string
	if exit == false {
		fmt.Println("[]============================================[]")
		fmt.Println("[]                  Isi Data                  []")
		fmt.Println("[]============================================[]")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 1. [] Isi Data Perusahaan                  []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 2. [] Isi Data Karyawan                    []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 0. [] Kembali                              []")
		fmt.Println("[]====[]======================================[]")
		fmt.Print("Masukkan inputan (0/2): ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "0" {
			fmt.Println("Masukkan inputan yang valid (0/2)")
			fmt.Scan(&opsi)
		}
		if opsi == "1" {
			isiDataPer(&P)
		} else if opsi == "2" {
			isiData(&K, &n)
		} else if opsi == "0" {
			menu()
		}
	}
}
