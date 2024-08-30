package main

import "fmt"

func editSaldoCuti(K *tabKaryawan, idx int) {
	var saldo_baru int
	if exit == false {
		fmt.Print("Ubah saldo cuti menjadi :")
		fmt.Scan(&saldo_baru)
		for saldo_baru < 0 || saldo_baru > 12 {
			fmt.Println("Masukkan inputan yang valid (0-12)")
			fmt.Scan(&saldo_baru)
		}
		K[idx].cuti.saldo = saldo_baru

		fmt.Println("[Saldo cuti karyawan berhasil diedit]")
		menuCuti(K, n)
	}
}
