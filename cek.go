package main

func cek(K tabKaryawan, cek_NIK string) bool {
	for i := 0; i < n; i++ {
		if K[i].NIK == cek_NIK {
			return false
		}
	}
	return true
}
