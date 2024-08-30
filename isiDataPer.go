package main

import "fmt"

func isiDataPer(P *perusahaan) {
	if exit == false {
		var nama, namaCeo, jenis, didirikan string

		fmt.Println("[]============================================[]")
		fmt.Print("Nama Perusahaan			: ")
		fmt.Scan(&nama)
		fmt.Print("Nama CEO Perusahaan		: ")
		fmt.Scan(&namaCeo)
		fmt.Print("Jenis Perusahaan		: ")
		fmt.Scan(&jenis)
		fmt.Print("Didirikan Pada Tahun		: ")
		fmt.Scan(&didirikan)
		fmt.Println("[]============================================[]")
		P.nama = nama
		P.ceo = namaCeo
		P.jenis = jenis
		P.didirikan = didirikan
		menuIsiData()
	}
}
