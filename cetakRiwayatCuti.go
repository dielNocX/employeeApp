package main

import "fmt"

func cetakRiwayatCuti(K *tabKaryawan, idx int) {
	//print cuti
	if exit == false {
		if K[idx].cuti.riwayat[0] == 0 {
			fmt.Println("[]=============================================================[]")
			fmt.Println("[]===========[Karyawan belum mempunyai riwayat cuti]===========[]")
			fmt.Println("[]=============================================================[]")
		} else {
			fmt.Println("[][Riwayat Cuti Karyawan]======================================[]")
			fmt.Println("[]=============================================================[]")
			fmt.Println("Nama		:", K[idx].nama)
			fmt.Println("Saldo Cuti	:", K[idx].cuti.saldo)

			for i := 0; i < (K)[idx].cuti.jumlah; i++ {
				fmt.Printf("%d. %d\n", i+1, K[idx].cuti.riwayat[i])
			}
		}
	}
}
