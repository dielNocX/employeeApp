package main

import "fmt"

func menuCuti(K *tabKaryawan, n int) {
	var opsi string
	// var indeks int
	var idx int = 0
	// var m int

	// cetakRiwayatCuti(idx)
	if exit == false {
		fmt.Println()
		fmt.Println("[]============================================[]")
		fmt.Println("[]                    CUTI                    []")
		fmt.Println("[]============================================[]")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 1. [] Cek Saldo Cuti Karyawan              []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 2. [] Edit Saldo Karyawan                  []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 3. [] Pengajuan Cuti Karyawan              []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 4. [] Tampilkan Riwayat Cuti Karyawan      []")
		fmt.Println("[]====[]======================================[]")
		fmt.Println("[] 0. [] Kembali                              []")
		fmt.Println("[]====[]======================================[]")
		fmt.Print("Masukkan inputan (0/4): ")
		fmt.Scan(&opsi)

		for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "0" {
			fmt.Println("Masukkan inputan yang valid (0/4) ")
			fmt.Scan(&opsi)
		}

		if opsi == "1" {
			cetakSaldoCuti(*K, n)
			menuCuti(K, n)

		} else if opsi == "2" {
			pilihKaryawan(*K, n, &idx)
			if idx >= 0 {
				editSaldoCuti(K, idx)
			}
			menuCuti(K, n)
		} else if opsi == "3" {
			pilihKaryawan(*K, n, &idx)
			if idx >= 0 {
				isiCuti(K, idx)
			}
			menuCuti(K, n)
		} else if opsi == "4" {
			pilihKaryawan(*K, n, &idx)
			if idx >= 0 {
				cetakRiwayatCuti(K, idx)
			}
			menuCuti(K, n)
		} else if opsi == "0" {
			menu()
		}

		// } else if opsi == "6" {
		// 	cutiKar(K, n)
	}
}
