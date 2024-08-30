package main

func descendingSort(K *tabKaryawan, n int, opsi int) {
	if exit == false {
		if opsi == 1 {
			for pass := 1; pass < n; pass++ {
				temp := K[pass]
				i := pass
				for i > 0 && temp.nama > K[i-1].nama {
					K[i] = K[i-1]
					i--
				}
				K[i] = temp
			}
		} else if opsi == 2 {
			for pass := 1; pass < n; pass++ {
				temp := K[pass]
				i := pass
				for i > 0 && temp.NIK > K[i-1].NIK {
					K[i] = K[i-1]
					i--
				}
				K[i] = temp
			}
		} else if opsi == 3 {
			for pass := 1; pass < n; pass++ {
				temp := K[pass]
				i := pass
				for i > 0 && temp.posisi > K[i-1].posisi {
					K[i] = K[i-1]
					i--
				}
				K[i] = temp
			}
		} else if opsi == 4 {
			for pass := 1; pass < n; pass++ {
				temp := K[pass]
				i := pass
				for i > 0 && temp.departemen > K[i-1].departemen {
					K[i] = K[i-1]
					i--
				}
				K[i] = temp
			}
		} else if opsi == 5 {
			for pass := 1; pass < n; pass++ {
				temp := K[pass]
				i := pass
				for i > 0 && temp.tglMasuk > K[i-1].tglMasuk {
					K[i] = K[i-1]
					i--
				}
				K[i] = temp
			}
		} else if opsi == 6 {
			for pass := 1; pass < n; pass++ {
				temp := K[pass]
				i := pass
				for i > 0 && temp.gender > K[i-1].gender {
					K[i] = K[i-1]
					i--
				}
				K[i] = temp
			}
		} else if opsi == 0 {
			menuDataTerurut()
		}
	}
}
