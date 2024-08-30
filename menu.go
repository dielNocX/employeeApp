package main

import (
	"fmt"
)

func menu() {

	var opsi string

	if exit == false {
		fmt.Println()
		fmt.Println("[]============================================[]")
		fmt.Println("[]                    MENU                    []")
		fmt.Println("[]============================================[]")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 1. [] Isi Data                             []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 2. [] Mencari Karyawan                     []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 3. [] Edit Data                            []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 4. [] Delete Data                          []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 5. [] Menampilkan Data                     []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 6. [] Ajukan cuti karyawan                 []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 0. [] Exit                                 []")
		fmt.Println("[]====[]======================================[]")
		fmt.Print("Masukkan inputan (0/6): ")

		fmt.Scan(&opsi)
		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" && opsi != "6" && opsi != "0" {
			fmt.Println("Masukkan inputan yang valid (0/6) ")
			fmt.Scan(&opsi)
		}

		if opsi == "0" {
			exit = true
			fmt.Print("exit :", exit)
		} else if opsi == "1" {
			fmt.Println()
			menuIsiData()
		} else if opsi == "2" {
			fmt.Println()
			cariData(K, n)
		} else if opsi == "3" {
			fmt.Println()
			editData(&K, n)
		} else if opsi == "4" {
			fmt.Println()
			deleteData(&K, &n)
		} else if opsi == "5" {
			fmt.Println()
			menuMenampilkanData()
		} else if opsi == "6" {
			fmt.Println()
			menuCuti(&K, n)
		}
	}
}
