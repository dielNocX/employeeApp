package main

const NMAX int = 1024

type perusahaan struct {
	nama string
	jenis string
	didirikan string
	ceo string
}

type cutiKaryawan struct {
	saldo   int
	riwayat [NMAX]int
	jumlah  int
}

type karyawan struct {
	nama, NIK, posisi, departemen, tglMasuk, gender string
	cuti                                            cutiKaryawan
}

type tabKaryawan [NMAX]karyawan

var K tabKaryawan
var n int
var P perusahaan
var exit bool

func main() {
	// isiDataPer(&P)
	menu()
}

// func menu() {
// 	var opsi string

// 	fmt.Println("Menu:")
// 	fmt.Println("1. Isi Data")
// 	fmt.Println("2. Mencari Karyawan")
// 	fmt.Println("3. Edit Data")
// 	fmt.Println("4. Delete Data")
// 	fmt.Println("5. Menampilkan Data")
// 	fmt.Println("6. Ajukan cuti karyawan")
// 	fmt.Println("Pilih opsi: ")

// 	fmt.Scan(&opsi)
// 	for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" && opsi != "6" {
// 		fmt.Println("Pilih 1/6: ")
// 		fmt.Scan(&opsi)
// 	}
// 	if opsi == "1" {
// 		isiData(&K, &n)
// 	} else if opsi == "2" {
// 		cariData(K, n)
// 	} else if opsi == "3" {
// 		editData(&K, n)
// 	} else if opsi == "4" {
// 		deleteData(&K, &n)
// 	} else if opsi == "5" {
// 		menampilkanData(K, n)
// 		// } else if opsi == "6" {
// 		// 	cutiKar(K, n)
// 	}
// }

// func isiData(K *tabKaryawan, n *int) {
// 	// IS: Array A dengan banyak elemen n terdefinisi sembarang
// 	// Proses: Membaca bilangan bulat dan memasukkan bilangan bulat positif ke dalam array A. Pembacaan dilakukan selama terbaca bilangan positif dan n < NMAX.
// 	// FS: Array A sebanyak n elemen berisi nilai
// 	var nama, NIK, posisi, departemen, tglMasuk, gender, konfirmasi string

// 	fmt.Print("Nama					: ")
// 	fmt.Scan(&nama)

// 	if nama != "." {
// 		fmt.Print("NIK					: ")
// 		fmt.Scan(&NIK)
// 		fmt.Print("Posisi					: ")
// 		fmt.Scan(&posisi)
// 		fmt.Print("Departemen				: ")
// 		fmt.Scan(&departemen)
// 		fmt.Print("Tanggal Masuk Karyawan			: ")
// 		fmt.Scan(&tglMasuk)
// 		fmt.Print("Gender					: ")
// 		fmt.Scan(&gender)
// 		if cek(*K, NIK) == false {
// 			fmt.Println("NIK telah terpakai, masukkan NIK yang berbeda")
// 			fmt.Print("NIK					: ")
// 			fmt.Scan(&NIK)
// 		}
// 	}
// 	for nama != "." && *n < NMAX {

// 		K[*n] = karyawan{
// 			nama:       nama,
// 			NIK:        NIK,
// 			posisi:     posisi,
// 			departemen: departemen,
// 			tglMasuk:   tglMasuk,
// 			gender:     gender,
// 			cuti:       cutiKaryawan{saldo: 12}}
// 		if *n < NMAX {
// 			*n++

// 		}
// 		fmt.Print("Ketik (.) untuk kembali ke menu atau ketik (Lanjutkan) untuk melanjutkan pengisian data: ")
// 		fmt.Scan(&konfirmasi)
// 		if konfirmasi != "." && konfirmasi == "Lanjutkan" {
// 			fmt.Print("Nama					: ")
// 			fmt.Scan(&nama)
// 			fmt.Print("NIK					: ")
// 			fmt.Scan(&NIK)
// 			fmt.Print("Posisi					: ")
// 			fmt.Scan(&posisi)
// 			fmt.Print("Departemen				: ")
// 			fmt.Scan(&departemen)
// 			fmt.Print("Tanggal Masuk Karyawan			: ")
// 			fmt.Scan(&tglMasuk)
// 			fmt.Print("Gender					: ")
// 			fmt.Scan(&gender)
// 			cariIndeks(*K, *n, 2, NIK)
// 			if cek(*K, NIK) == false {
// 				fmt.Println("NIK telah terpakai, masukkan NIK yang berbeda")
// 				fmt.Print("NIK					: ")
// 				fmt.Scan(&NIK)
// 			}
// 		} else {
// 			menu()
// 		}

// 	}
// }

// // func mencariData() {
// // 	var opsi int
// // 	for opsi != 1 || opsi != 2 {
// // 		fmt.Println("Cari data berdasarkan: ")
// // 		fmt.Println("1. Nama")
// // 		fmt.Println("2. NIK")
// // 		fmt.Println("3. Kembali")
// // 	}

// // 	if opsi == 1 {
// // 		mencariDataNama()
// // 	} else if opsi == 2 {
// // 		mencarDataNIK()
// // 	} else if opsi == 3 {
// // 		menu()
// // 	}
// // }

// // func mencariDataNama() {

// // }

// // func mencariDataNIK() {

// // }

// // func editData() {

// // }

// func menampilkanData(k tabKaryawan, n int) {
// 	for i := 0; i < n; i++ {
// 		fmt.Println("[]=============================================================[]")
// 		fmt.Println("[][Data karyawan ke-", i+1, "]=======================================[]")
// 		fmt.Println("[]=============================================================[]")
// 		fmt.Printf("Nama					: %s\n", k[i].nama)
// 		fmt.Printf("NIK					: %s\n", k[i].NIK)
// 		fmt.Printf("Posisi					: %s\n", k[i].posisi)
// 		fmt.Printf("Departemen				: %s\n", k[i].departemen)
// 		fmt.Printf("Tanggal Masuk Karyawan			: %s\n", k[i].tglMasuk)
// 		fmt.Printf("Gender					: %s\n", k[i].gender)
// 		if i == n-1 {
// 			fmt.Println("[]=============================================================[]")
// 		}
// 	}
// 	if k[0].nama == "" {
// 		fmt.Println("[]=============================================================[]")
// 		fmt.Println("[]===============[DATA KARYAWAN TIDAK DITEMUKAN]===============[]")
// 		fmt.Println("[]=============================================================[]")
// 		menu()
// 	}
// 	menu()
// }

// func editData(K *tabKaryawan, n int) {
// 	// IS: x adalah input yg dicari pada tabel dan y adalah opsi (1,2,3) untuk menentukan parameter yang dicari
// 	// FS:
// 	var indeks int
// 	var opsi, nama_baru, NIK_baru, posisi_baru, departemen_baru, tglMasuk_baru, gender_baru, cari_NIK string

// 	fmt.Print("Masukkan data NIK karyawan yang ingin anda edit: ")
// 	fmt.Scan(&cari_NIK)
// 	indeks = cariIndeks(*K, n, 2, cari_NIK)
// 	cetak(indeks)

// 	fmt.Println("Apa yang ingin anda edit: ")
// 	fmt.Println("1. Nama")
// 	fmt.Println("2. NIK")
// 	fmt.Println("3. Posisi")
// 	fmt.Println("4. Departemen")
// 	fmt.Println("5. Tanggal Masuk Karyawan")
// 	fmt.Println("6. Gender")
// 	fmt.Scan(&opsi)

// 	for opsi != "1" && opsi != "2" && opsi != "3" && opsi != "4" && opsi != "5" && opsi != "6" {
// 		fmt.Println("Masukkan inputan yang valid pilih (1/6): ")
// 		fmt.Scan(&opsi)
// 	}
// 	if opsi == "1" {
// 		fmt.Print("Ganti Nama menjadi: ")
// 		fmt.Scan(&nama_baru)
// 		K[indeks].nama = nama_baru
// 	} else if opsi == "2" {
// 		fmt.Print("\nGanti NIK menjadi: ")
// 		fmt.Scan(&NIK_baru)
// 		K[indeks].NIK = NIK_baru
// 	} else if opsi == "3" {
// 		fmt.Print("\nGanti Posisi menjadi: ")
// 		fmt.Scan(&posisi_baru)
// 		K[indeks].posisi = posisi_baru
// 	} else if opsi == "4" {
// 		fmt.Print("\nGanti Departemen menjadi: ")
// 		fmt.Scan(&departemen_baru)
// 	} else if opsi == "5" {
// 		fmt.Print("\nGanti Tanggal Masuk Karyawan Menjadi: ")
// 		fmt.Scan(&tglMasuk_baru)
// 	} else if opsi == "6" {
// 		fmt.Print("\nGanti Gender menjadi: ")
// 		fmt.Scan(&gender_baru)
// 	}
// 	cetak(indeks)
// 	menu()
// }

// func cariData(K tabKaryawan, n int) {
// 	var opsi, x string
// 	var indeks int
// 	fmt.Println("Cari data karyawan berdasarkan: ")
// 	fmt.Println("1. Nama")
// 	fmt.Println("2. NIK")
// 	fmt.Println("3. Posisi")
// 	fmt.Println("4. Departemen")
// 	fmt.Println("5. Tanggal Masuk Karyawan")
// 	fmt.Println("6. Gender")
// 	fmt.Scan(&opsi)
// 	if opsi == "1" {
// 		fmt.Println("Masukkan data nama karyawan yang ingin anda cari: ")
// 		fmt.Scan(&x)
// 		indeks = cariIndeks(K, n, 1, x)
// 	} else if opsi == "2" {
// 		fmt.Println("Masukkan data NIK karyawan yang ingin anda cari: ")
// 		fmt.Scan(&x)
// 		indeks = cariIndeks(K, n, 2, x)
// 	} else if opsi == "3" {
// 		fmt.Println("Masukkan data Posisi karyawan yang ingin anda cari: ")
// 		fmt.Scan(&x)
// 		indeks = cariIndeks(K, n, 3, x)
// 	} else if opsi == "4" {
// 		fmt.Println("Masukkan data Departemen karyawan yang ingin anda cari: ")
// 		fmt.Scan(&x)
// 		indeks = cariIndeks(K, n, 4, x)
// 	} else if opsi == "5" {
// 		fmt.Println("Masukkan data Tanggal Masuk Karyawan yang ingin anda cari: ")
// 		fmt.Scan(&x)
// 		indeks = cariIndeks(K, n, 5, x)
// 	} else if opsi == "6" {
// 		fmt.Println("Masukkan data Gender karyawan yang ingin anda cari: ")
// 		fmt.Scan(&x)
// 		indeks = cariIndeks(K, n, 6, x)
// 	} else {
// 		fmt.Println("[]=============================================================[]")
// 		fmt.Println("[]===============[DATA KARYAWAN TIDAK DITEMUKAN]===============[]")
// 		fmt.Println("[]=============================================================[]")
// 	}
// 	cetak(indeks)
// 	menu()
// }

// func cariIndeks(K tabKaryawan, n, opsi int, x string) int {
// 	var indeks int
// 	for i := 0; i < n; i++ {
// 		if opsi == 1 {
// 			if K[i].nama == x {
// 				indeks = i
// 			}
// 		} else if opsi == 2 {
// 			if K[i].NIK == x {
// 				indeks = i
// 			}
// 		} else if opsi == 3 {
// 			if K[i].posisi == x {
// 				indeks = i
// 			}
// 		} else if opsi == 4 {
// 			if K[i].departemen == x {
// 				indeks = i
// 			}
// 		} else if opsi == 5 {
// 			if K[i].tglMasuk == x {
// 				indeks = i
// 			}
// 		} else if opsi == 6 {
// 			if K[i].gender == x {
// 				indeks = i
// 			}
// 		}
// 	}
// 	return indeks
// }

// func deleteData(K *tabKaryawan, n *int) {
// 	var nama, opsi string
// 	var delete karyawan
// 	var indeks int
// 	fmt.Println("Masukkan nama yang ingin anda delete: ")
// 	fmt.Scan(&nama)
// 	indeks = cariIndeks(*K, *n, 1, nama)
// 	cetak(indeks)
// 	fmt.Println("Apakah anda yakin ingin delete data karyawan berikut? ")
// 	fmt.Scan(&opsi)
// 	if opsi == "yes" {
// 		fmt.Print(indeks)
// 		for indeks < *n-1 {
// 			K[indeks] = K[indeks+1]
// 			indeks++
// 		}
// 		K[indeks] = delete
// 		*n -= 1
// 	}
// 	menu()
// }

// func cetak(i int) {
// 	fmt.Println("[]=============================================================[]")
// 	fmt.Println("[][Data karyawan ke-", i+1, "]=======================================[]")
// 	fmt.Println("[]=============================================================[]")
// 	fmt.Printf("Nama					: %s\n", K[i].nama)
// 	fmt.Printf("NIK					: %s\n", K[i].NIK)
// 	fmt.Printf("Posisi					: %s\n", K[i].posisi)
// 	fmt.Printf("Departemen				: %s\n", K[i].departemen)
// 	fmt.Printf("Tanggal Masuk Karyawan			: %s\n", K[i].tglMasuk)
// 	fmt.Printf("Gender					: %s\n", K[i].gender)
// 	fmt.Println("[]=============================================================[]")
// }

// // func cutiKar(K tabKaryawan, n int) {
// // 	var NIK_karyawan string
// // 	var indeks, i, j int
// // 	fmt.Println("Masukkan nama karyawan: ")
// // 	fmt.Scan(&NIK_karyawan)
// // 	indeks = cariIndeks(K, n, 2, NIK_karyawan)
// // 	cetak(indeks)
// // 	fmt.Println("Masukkan tanggal cuti karyawan: ")
// // 	fmt.Scan(&K[i].cuti.riwayat[j])
// // 	K[indeks].cuti.saldo += 1
// // 	if K[indeks].cuti.saldo > 12 {
// // 		K[indeks].cuti.saldo = 12
// // 	}
// // 	for j < K[indeks].cuti.saldo {
// // 		fmt.Println(K[indeks].cuti.riwayat[j])
// // 		j++
// // 	}
// // }

// func cek(K tabKaryawan, cek_NIK string) bool {
// 	for i := 0; i < n; i++ {
// 		if K[i].NIK == cek_NIK {
// 			return false
// 		}
// 	}
// 	return true
// }
