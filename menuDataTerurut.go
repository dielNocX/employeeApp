package main

import "fmt"

func menuDataTerurut() {
	var opsi string
	if exit == false {
		fmt.Println("[]============================================[]")
		fmt.Println("[]           Tampilkan Data Terurut           []")
		fmt.Println("[]============================================[]")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 1. [] Ascending                            []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 2. [] Descending                           []")
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
			ascending(&K, n)
		} else if opsi == "2" {
			descending(&K, n)
		} else if opsi == "0" {
			menuMenampilkanData()
		}
	}
}
