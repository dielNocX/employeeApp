package main

func ascendingSort(K *tabKaryawan, n int, opsi int) {
	var temp karyawan
	var pass, idx, i int
	if exit == false {
		if opsi == 1 {
			for pass = 0; pass < n; pass++ {
				idx = pass
				i = pass + 1
				temp = (*K)[pass]
				for i < n {
					if K[idx].nama > K[i].nama {
						idx = i
					}
					i++
				}
				temp = K[pass]
				K[pass] = K[idx]
				K[idx] = temp
			}
		} else if opsi == 2 {
			for pass = 0; pass < n; pass++ {
				idx = pass
				i = pass + 1
				temp = (*K)[pass]
				for i < n {
					if K[idx].NIK > K[i].NIK {
						idx = i
					}
					i++
				}
				temp = K[pass]
				K[pass] = K[idx]
				K[idx] = temp
			}
		} else if opsi == 3 {
			for pass = 0; pass < n; pass++ {
				idx = pass
				i = pass + 1
				temp = (*K)[pass]
				for i < n {
					if K[idx].posisi > K[i].posisi {
						idx = i
					}
					i++
				}
				temp = K[pass]
				K[pass] = K[idx]
				K[idx] = temp
			}
		} else if opsi == 4 {
			for pass = 0; pass < n; pass++ {
				idx = pass
				i = pass + 1
				temp = (*K)[pass]
				for i < n {
					if K[idx].departemen > K[i].departemen {
						idx = i
					}
					i++
				}
				temp = K[pass]
				K[pass] = K[idx]
				K[idx] = temp
			}
		} else if opsi == 5 {
			for pass = 0; pass < n; pass++ {
				idx = pass
				i = pass + 1
				temp = (*K)[pass]
				for i < n {
					if K[idx].tglMasuk > K[i].tglMasuk {
						idx = i
					}
					i++
				}
				temp = K[pass]
				K[pass] = K[idx]
				K[idx] = temp
			}
		} else if opsi == 6 {
			for pass = 0; pass < n; pass++ {
				idx = pass
				i = pass + 1
				temp = (*K)[pass]
				for i < n {
					if K[idx].gender > K[i].gender {
						idx = i
					}
					i++
				}
				temp = K[pass]
				K[pass] = K[idx]
				K[idx] = temp
			}
		} else if opsi == 0 {
			menuDataTerurut()
		}
	}
}
