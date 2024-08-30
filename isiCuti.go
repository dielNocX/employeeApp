package main

import "fmt"

func isiCuti(K *tabKaryawan, idx int) {
	//mengajukan Cuti
	var tglCuti, nCuti int
	if exit == false {
		if *&K[idx].cuti.saldo > 0 {
			fmt.Print("Berapa hari cuti yang diajukan : ")
			fmt.Scan(&nCuti)
			for nCuti > K[idx].cuti.saldo {
				fmt.Println("[]=============================================================[]")
				fmt.Println("[][Karyawan tidak mempunyai saldo cuti yang cukup]=============[]")
				fmt.Printf("Sisa saldo	: %d\n", K[idx].cuti.saldo)
				fmt.Printf("[]=============================================================[]\n")
				fmt.Print("Berapa hari cuti yang diajukan : ")
				fmt.Scan(&nCuti)
			}
			fmt.Println("Masukkan tanggal cuti karyawan dengan format DDMMYYYY")
			for i := (*K)[idx].cuti.jumlah; i < (*K)[idx].cuti.jumlah+nCuti; i++ {
				fmt.Scan(&tglCuti)
				for tglCuti < 1 {
					fmt.Println("[Masukkan inputan yang valid]")
					fmt.Scan(&tglCuti)
				}
				// if cekCuti(*K, idx, *m, tglCuti) == false {
				(*K)[idx].cuti.riwayat[i] = tglCuti
				// fmt.Println((K)[idx].cuti.riwayat[i])
				// }
				(*K)[idx].cuti.saldo--
				tglCuti = 0
			}
			(*K)[idx].cuti.jumlah = (*K)[idx].cuti.jumlah + nCuti
		} else {
			fmt.Println("[]============[Karyawan tidak mempunyai saldo cuti]============[]")
			fmt.Println("[]=============================================================[]")
		}
		menuCuti(K, n)
	}
}
