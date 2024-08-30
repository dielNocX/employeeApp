package main

func cariIndeks(K tabKaryawan, n, opsi int, x string) int {
	var indeks int = -1
	for i := 0; i < n; i++ {
		if opsi == 1 {
			if K[i].nama == x {
				indeks = i
			}
		} else if opsi == 2 {
			if K[i].NIK == x {
				indeks = i
			}
		} else if opsi == 3 {
			if K[i].posisi == x {
				indeks = i
			}
		} else if opsi == 4 {
			if K[i].departemen == x {
				indeks = i
			}
		} else if opsi == 5 {
			if K[i].tglMasuk == x {
				indeks = i
			}
		} else if opsi == 6 {
			if K[i].gender == x {
				indeks = i
			}
		}
	}
	return indeks
}
