package main

func cekCuti(K tabKaryawan, idx, x int) bool {

	for i := 0; i < (K)[idx].cuti.jumlah; i++ {

		if K[idx].cuti.riwayat[i] == x {
			return true
		}
	}
	return false
}
