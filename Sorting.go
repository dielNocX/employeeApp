package main

func selectionSort(K *tabKaryawan, n int) {

	var pass, k, idx_min int
	if exit == false {
		pass = 0
		for pass < n-1 {
			idx_min = pass
			k = pass + 1
			for k < n {
				if K[idx_min].nama < K[k].nama {
					idx_min = k
				}
				k++
			}

			K[pass], K[idx_min] = K[idx_min], K[pass]
			pass++
		}
		menuMenampilkanData()
	}
}
