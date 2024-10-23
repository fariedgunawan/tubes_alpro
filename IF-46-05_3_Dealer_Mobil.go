package main

import "fmt"

const NMAX int = 2023

type dealer struct {
	nama     string
	pabrikan string
	warna    string
	tahun    int
	kode     int
}

type pabrikan struct {
	nama      string
	jumlahnya int
}

type arrPabrikan [100]pabrikan

type arrDealer [NMAX]dealer

func main() {
	var arr arrDealer
	var arrPab arrPabrikan
	var i int
	var opsi int
	var pilihan string

	fmt.Println("")
	fmt.Println("Selamat Datang di Aplikasi Management Mobil:")
	fmt.Println("")

	Header()
	fmt.Println("Masukan Opsi Dibawah ini:")
	fmt.Scan(&opsi)

	for opsi != 8 {
		if opsi == 1 {
			inputData(&arr, &i, &arrPab)
			fmt.Println("====== Masukan Selesai =======") // input data
		} else if opsi == 2 {
			printData(arr, i) // print data
		} else if opsi == 3 {
			searchData(arr, i) // search data
		} else if opsi == 4 {
			editData(&arr, i, &arrPab) // edit data
		} else if opsi == 5 {
			inputDataAgain(&arr, &i, &arrPab) // input data kembali
			fmt.Println("====== Masukan Sudah Terupdate =======")
		} else if opsi == 6 {
			deleteData(&arr, &i, &arrPab) // delete data
		} else if opsi == 7 {
			SelectionSort(&arrPab, i)
			printTop3(&arrPab) // menampilkan top 3
		} else {
			fmt.Println("Opsi Tidak Ada")
			fmt.Println("")
		}

		fmt.Println("Tampilkan opsi kembali?")
		fmt.Println("yes / no")
		fmt.Scan(&pilihan)

		for pilihan != "yes" && pilihan != "no" {
			fmt.Println("")
			fmt.Println("Opsi Tidak Tersedia")
			fmt.Scan(&pilihan)
		}
		

		if pilihan == "yes" {
			Header()
			fmt.Println("Masukan Kembali Opsi Dibawah ini:")
			fmt.Scan(&opsi)
		} else if pilihan == "no" {
			fmt.Println("")
			fmt.Println("Masukan Opsi dibawah ini:")
			fmt.Scan(&opsi)
		}
	}
	fmt.Println("====== Program Selesai =========")
}

func inputData(arr *arrDealer, i *int, arrPab *arrPabrikan) {
	var pabrikan, nama, warna string
	var kode, tahun int

	fmt.Println("Masukan Data")
	fmt.Println("==================")
	fmt.Scan(&pabrikan)

	for pabrikan != "selesai" {
		fmt.Scan(&nama, &warna, &tahun, &kode)

		arr[*i].pabrikan = pabrikan
		arr[*i].nama = nama
		arr[*i].warna = warna
		arr[*i].tahun = tahun
		arr[*i].kode = kode

		foundIndex := -1 

		for c := 0; c < *i && foundIndex == -1; c++ {  // memasukan data kedalam tabel
			if pabrikan == arrPab[c].nama {
				foundIndex = c
			}
		}

		if foundIndex != -1 {
			arrPab[foundIndex].jumlahnya++
		} else {
			arrPab[*i].nama = pabrikan
			arrPab[*i].jumlahnya++
		}
		*i++
		fmt.Scan(&pabrikan)
	}
}

func Header() {
	fmt.Println("===========================")
	fmt.Println("|    Terdapat opsi        |")
	fmt.Println("|-------------------------|")
	fmt.Println("| 1. Input Data           |")
	fmt.Println("| 2. Tampilkan Data       |")
	fmt.Println("| 3. Search Data          |")
	fmt.Println("| 4. Forum Edit Data      |")
	fmt.Println("| 5. Input Data Kembali   |")
	fmt.Println("| 6. Delete Data          |")
	fmt.Println("| 7. Penjualan Terbanyak  |")
	fmt.Println("| 8. Quit                 |")
	fmt.Println("===========================")
}

func inputDataAgain(arr *arrDealer, startIdx *int, arrPab *arrPabrikan) {
	var pabrikan, nama, warna string
	var kode, tahun int

	fmt.Println("")
	fmt.Println("Masukan Data")
	fmt.Println("==================")
	fmt.Scan(&pabrikan)

	for pabrikan != "selesai" {
		fmt.Scan(&nama, &warna, &tahun, &kode)
		arr[*startIdx].pabrikan = pabrikan
		arr[*startIdx].nama = nama
		arr[*startIdx].warna = warna
		arr[*startIdx].tahun = tahun
		arr[*startIdx].kode = kode

		foundIndex := -1

		for c := 0; c < *startIdx && foundIndex == -1; c++ { // memasukan data kedalam tabel
			if pabrikan == arrPab[c].nama {
				foundIndex = c
			}
		}

		if foundIndex != -1 {
			arrPab[foundIndex].jumlahnya++
		} else {
			arrPab[*startIdx].nama = pabrikan
			arrPab[*startIdx].jumlahnya++
		}
		*startIdx++
		fmt.Scan(&pabrikan)
	}
}

func printData(arr arrDealer, size int) {
	fmt.Println("================")
	fmt.Println("Data dalam array")
	fmt.Println("================")

	for b := 0; b < size; b++ {
		fmt.Printf("Nama Pabrikan: %v\t Nama mobil: %v\t Warna mobil: %v\t Tahun: %d\t Kode mobil: %d\n", arr[b].pabrikan, arr[b].nama, arr[b].warna, arr[b].tahun, arr[b].kode)
	}
	fmt.Println("")

}

func searchData(arr arrDealer, size int) {
	var cari, pabrikan, nama, warna string
	var kode, tahun int

	fmt.Println("==============================")
	fmt.Println("     Opsi yang tersedia:")
	fmt.Println("==============================")
	fmt.Println("pabrikan")
	fmt.Println("nama")
	fmt.Println("warna")
	fmt.Println("tahun")
	fmt.Println("kode")
	fmt.Println("==============================")
	fmt.Println("Masukkan opsi yang ingin dicari:")
	fmt.Scan(&cari)

	if cari == "pabrikan" {
		fmt.Println("Masukan Nama Pabrikannya")
		fmt.Scan(&pabrikan)
		fmt.Println("======  Hasil  =======")
		for c := 0; c < size; c++ {
			if arr[c].pabrikan == pabrikan {
				fmt.Println("Nama Mobil tersebut:", arr[c].nama, "Dengan warna:", arr[c].warna, "Tahun:", arr[c].tahun, "Kode:", arr[c].kode)
			}
		}
	} else if cari == "nama" {
		fmt.Println("Masukan Nama Mobilnya")
		fmt.Scan(&nama)
		fmt.Println("======Hasil=======")
		for c := 0; c < size; c++ {
			if arr[c].nama == nama {
				fmt.Println("Pabrikan Mobil tersebut:", arr[c].pabrikan, "Dengan warna:", arr[c].warna, "Tahun:", arr[c].tahun, "Kode:", arr[c].kode)
			}
		}
	} else if cari == "warna" {
		fmt.Println("Masukan Warna Mobilnya")
		fmt.Scan(&warna)
		fmt.Println("======Hasil=======")
		for c := 0; c < size; c++ {
			if arr[c].warna == warna {
				fmt.Println("Pabrikan Mobil tersebut:", arr[c].pabrikan, "Dengan Nama:", arr[c].nama, "Tahun:", arr[c].tahun, "Kode:", arr[c].kode)
			}
		}
	} else if cari == "tahun" {
		fmt.Println("Masukan Tahun Keluaran Mobilnya")
		fmt.Scan(&tahun)
		fmt.Println("======Hasil=======")
		for c := 0; c < size; c++ {
			if arr[c].tahun == tahun {
				fmt.Println("Pabrikan Mobil tersebut:", arr[c].pabrikan, "Dengan Nama:", arr[c].nama, "Warna Mobil Tersebut:", arr[c].warna, "Kode:", arr[c].kode)
			}
		}
	} else if cari == "kode" {
		fmt.Println("Masukan Kode Mobilnya")
		fmt.Scan(&kode)
		fmt.Println("======Hasil=======")
		for c := 0; c < size; c++ {
			if arr[c].kode == kode {
				fmt.Println("Pabrikan Mobil tersebut:", arr[c].pabrikan, "Dengan Nama:", arr[c].nama, "Tahun:", arr[c].tahun, "Warna Mobil Tersebut:", arr[c].warna)
			}
		}
	}
	fmt.Println("")
}

func editData(arr *arrDealer, size int, arrPab *arrPabrikan) {
	var kode, tahun, j int
	var cari, pabrikan, nama, warna string

	fmt.Println("=================")
	fmt.Println("Forum Pengeditan")
	fmt.Println("=================")

	fmt.Println("Masukan Kode Mobil:")
	fmt.Scan(&kode)
	fmt.Println("")

	for j = 0; j < size; j++ {
		if arr[j].kode == kode {
			fmt.Println("Masukan nama variabel:")
			fmt.Scan(&cari)
			if cari == "pabrikan" {
				fmt.Println("Masukan Perubahannya:")
				fmt.Scan(&pabrikan)
				arr[j].pabrikan = pabrikan
			} else if cari == "nama" {
				fmt.Println("Masukan Perubahannya:")
				fmt.Scan(&nama)
				arr[j].nama = nama
			} else if cari == "warna" {
				fmt.Println("Masukan Perubahannya:")
				fmt.Scan(&warna)
				arr[j].warna = warna
			} else if cari == "tahun" {
				fmt.Println("Masukan Perubahannya:")
				fmt.Scan(&tahun)
				arr[j].tahun = tahun
			}
		}
	}

	fmt.Println("")
	fmt.Println("=========== HASIL SUDAH EDIT =============")
	fmt.Println("")
	// data array kembali dipanggil
	for j := 0; j < size; j++ {
		fmt.Printf("Nama Pabrikan: %v\t Nama mobil: %v\t Warna mobil: %v\t Tahun: %d\t Kode mobil: %d\n", arr[j].pabrikan, arr[j].nama, arr[j].warna, arr[j].tahun, arr[j].kode)
	}
	fmt.Println("")
}

func seqSeacrh(arr arrDealer, size, kode int) int {
	var i int
	for i < size && kode != arr[i].kode {
		i++
	}
	if i == size {
		return -1
	} else {
		return i
	}
}

func deleteData(arr *arrDealer, size *int, arrPab *arrPabrikan) {
	var kode, j, idx int

	fmt.Println("=================")
	fmt.Println("Hapus Data")
	fmt.Println("=================")

	fmt.Println("Masukan Kode Mobil:")
	fmt.Scan(&kode)
	fmt.Println("")
	idx = seqSeacrh(*arr, *size, kode)
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
	
		foundIndex := -1
		for c := 0; c < *size && foundIndex == -1; c++ {
			if arrPab[c].nama == arr[idx].pabrikan {
				foundIndex = c
			}
		}
		if foundIndex != -1 {
			arrPab[foundIndex].jumlahnya--
		}

		for j = idx; j < *size-1; j++ {
			arr[j] = arr[j+1]
		}
		*size--

		fmt.Println("Data berhasil dihapus")
	}
}

func MenampilkanTabel(arrPab arrPabrikan, size int) {
	for s := 0; s < size; s++ {
		if arrPab[s].nama != "" {
			fmt.Printf("Nama: %s\t Jumlah: %d\n", arrPab[s].nama, arrPab[s].jumlahnya)
		}
	}
}

func SelectionSort(arrPab *arrPabrikan, size int) {
	var (
		i, j, maxIndex int
		temp           pabrikan
	)
	for i = 0; i < size-1; i++ {
		maxIndex = i
		for j = i + 1; j < size; j++ {
			if arrPab[j].jumlahnya > arrPab[maxIndex].jumlahnya {
				maxIndex = j
			}
		}
		if maxIndex != i {
			temp = arrPab[maxIndex]
			arrPab[maxIndex] = arrPab[i]
			arrPab[i] = temp
		}
	}
}

func printTop3(arrPab *arrPabrikan) {
	fmt.Println("Pabrikan dengan penjualan terbanyak:")
	for i := 0; i < 3; i++ {
		if arrPab[i].nama != "" {
			fmt.Printf("Peringkat %d: Pabrikan %s dengan jumlah %d penjualan\n", i+1, arrPab[i].nama, arrPab[i].jumlahnya)
		}
	}
}

// contoh input
// honda jazz hitam 2020 1
// honda yaris merah 2020 2
// suzuki ertiga hitam 2020 3
// ferari spider merah 2020 4
// bugati chiron hitam 2020 5
// honda freed merah 2020 6
// honda crv hitam 2020 7
// suzuki papa merah 2020 8
// suzuki pipi merah 2020 9
// bugati pap hitam 2020 10
// selesai 
