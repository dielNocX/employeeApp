package main

import "fmt"

func menampilkanDataPer(P *perusahaan) {
	if exit == false {
		if P.nama != "" {
			fmt.Println("[]============================================[]")
			fmt.Printf("Nama Perusahaan			: %s\n", P.nama)
			fmt.Printf("Nama CEO Perusahaan		: %s\n", P.ceo)
			fmt.Printf("Jenis Perusahaan		: %s\n", P.jenis)
			fmt.Printf("Didirikan Pada Tahun		: %s\n", P.didirikan)
			fmt.Printf("Jumlah Karyawan			: %d\n", n)
			fmt.Println("[]============================================[]")
		} else {
			fmt.Println("[]=============================================================[]")
			fmt.Println("[]==============[DATA PERUSAHAAN TIDAK DITEMUKAN]==============[]")
			fmt.Println("[]=============================================================[]")
		}
		menu()
	}
}
