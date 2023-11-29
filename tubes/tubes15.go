package main

import (
	"fmt"
	"strconv"
)

type tempatWisata struct {
	nama, lokasi, kontak, hari, waktuOps string
	harga                                int
}

const NMAX int = 100

type tabWisata [NMAX]tempatWisata

func main() {
	var daftarWisata tabWisata
	var n int
	//dummy data
	daftarWisata[0] = tempatWisata{"Dufan", "Jakarta", "081212", "Senin-Sabtu", "07.00-19.00", 15000}
	daftarWisata[1] = tempatWisata{"Ancol", "Jakarta", "046285", "Selasa-Sabtu", "08.00-20.00", 24000}
	daftarWisata[2] = tempatWisata{"Puncak", "Bogor", "573967", "Rabu-Kamis", "07.30-16.00", 10000}
	daftarWisata[3] = tempatWisata{"Lembang", "Bandung", "23649", "Sabtu-Minggu", "09.00-18.30", 10000}
	daftarWisata[4] = tempatWisata{"Gunung_Padang", "Cianjur", "748903", "Senin-Rabu", "10.00-21.00", 18000}
	daftarWisata[5] = tempatWisata{"Taman_Safari", "Bogor", "583659", "Selasa-Jumat", "08.30-15.30", 23000}
	daftarWisata[6] = tempatWisata{"Selabintana", "Sukabumi", "894633", "Rabu-Minggu", "07.00-20.00", 29000}
	daftarWisata[7] = tempatWisata{"Pulau_Seribu", "Jakarta", "563859", "Kamis-Minggu", "09.00-17.30", 17000}
	daftarWisata[8] = tempatWisata{"Pelabuhan_Ratu", "Sukabumi", "362057", "Rabu-Minggu", "10.00-19.30", 18500}
	daftarWisata[9] = tempatWisata{"Kawah_Talaga_Bodas", "Garut", "890146", "Kamis-Jumat", "07.00-17.00", 21000}
	daftarWisata[10] = tempatWisata{"Kebun_Raya_Bogor", "Bogor", "174294", "Selasa-Minggu", "07.30-16.30", 23000}
	n = 11
	mainMenu(&daftarWisata, &n)
}

// Input data
func inputArrWisata(T *tabWisata, n *int) {
	var nama, lokasi, hari, kontak, waktuOps string
	var harga int

	header()
	fmt.Println("Menginput data tempat wisata...")
	fmt.Print("Nama tempat\t: ")
	fmt.Scan(&nama)
	if findNama(*T, *n, nama) != -1 {
		fmt.Println("Tempat wisata", nama, "sudah ada")
	} else {
		fmt.Print("Lokasi kota\t: ")
		fmt.Scan(&lokasi)
		fmt.Print("Harga tiket\t: ")
		fmt.Scan(&harga)
		fmt.Print("Hari buka\t: ")
		fmt.Scan(&hari)
		fmt.Print("Jam operasional\t: ")
		fmt.Scan(&waktuOps)
		fmt.Print("Kontak\t\t: ")
		fmt.Scan(&kontak)
		if konfirmasi("Simpan data?") {
			T[*n] = tempatWisata{nama, lokasi, kontak, hari, waktuOps, harga}
			*n++
			fmt.Println("Data berhasil di simpan")
		}
	}
}

// Mencetak data wisata dengan tabel
func printArrwisata(T tabWisata, n int) {
	var row tempatWisata
	var lebar [6]int
	for i := 0; i < n; i++ {
		row = T[i]
		if len(row.nama) > lebar[0] {
			lebar[0] = len(row.nama)
		}
		if len(row.lokasi) > lebar[1] {
			lebar[1] = len(row.lokasi)
		}
		if len(strconv.Itoa(row.harga)) > lebar[2] {
			lebar[2] = len(strconv.Itoa(row.harga))
		}
		if len(row.hari) > lebar[3] {
			lebar[3] = len(row.hari)
		}
		if len(row.waktuOps) > lebar[4] {
			lebar[4] = len(row.waktuOps)
		}
		if len(row.kontak) > lebar[5] {
			lebar[5] = len(row.kontak)
		}
	}

	var digitN int = len(strconv.Itoa(n)) + 1
	var pemisah string = "+--"
	for i := 0; i < digitN; i++ {
		pemisah += "-"
	}
	for i := 0; i < len(lebar); i++ {
		pemisah += "+"
		for j := 0; j < lebar[i]+2; j++ {
			pemisah += "-"
		}
	}
	pemisah += "+"

	fmt.Println(pemisah)
	fmt.Printf("| %-*s ", digitN, "No")
	fmt.Printf("| %-*s ", lebar[0], "Nama")
	fmt.Printf("| %-*s ", lebar[1], "Kota")
	fmt.Printf("| %-*s ", lebar[2], "Harga")
	fmt.Printf("| %-*s ", lebar[3], "Hari Buka")
	fmt.Printf("| %-*s ", lebar[4], "Jam Buka")
	fmt.Printf("| %-*s |\n", lebar[5], "Kontak")
	for i := 0; i < n; i++ {
		fmt.Println(pemisah)
		row = T[i]
		fmt.Printf("| %-*d ", digitN, i+1)
		fmt.Printf("| %-*s ", lebar[0], row.nama)
		fmt.Printf("| %-*s ", lebar[1], row.lokasi)
		fmt.Printf("| %-*d ", lebar[2], row.harga)
		fmt.Printf("| %-*s ", lebar[3], row.hari)
		fmt.Printf("| %-*s ", lebar[4], row.waktuOps)
		fmt.Printf("| %-*s |\n", lebar[5], row.kontak)
	}
	fmt.Println(pemisah)
}

// Mengurutkan harga dengan selection sort
func urutDataHarga(T *tabWisata, n int, filter string) {
	var idx int
	var i, j int
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = i
		for j < n {
			if (filter == "naik" && T[idx].harga > T[j].harga) ||
				(filter == "turun" && T[idx].harga < T[j].harga) {
				idx = j
			}
			j++
		}
		T[idx], T[i-1] = T[i-1], T[idx]
		i++
	}

}

// Mengurutkan nama dengan insertion sort
func urutDataNama(T *tabWisata, n int, filter string) {
	var temp tempatWisata
	var i, j int
	for i = 1; i < n; i++ {
		j = i
		temp = T[j]
		if filter == "naik" {
			for j > 0 && temp.nama < T[j-1].nama {
				T[j] = T[j-1]
				j--
			}
		} else if filter == "turun" {
			for j > 0 && temp.nama > T[j-1].nama {
				T[j] = T[j-1]
				j--
			}
		}
		T[j] = temp
	}
}

// Mencari index berdasarkan nama dengan binary search
func findNama(T tabWisata, n int, keyword string) int {
	var kiri, kanan, tengah, idx int

	idx = -1
	kiri = 0
	kanan = n - 1
	for kiri <= kanan && idx == -1 {
		tengah = (kiri + kanan) / 2
		if T[tengah].nama < keyword {
			kiri = tengah + 1
		} else if T[tengah].nama > keyword {
			kanan = tengah - 1
		} else {
			idx = tengah
		}
	}
	return idx
}

// Mencetak data berdasarkan lokasi yang ditemukan dengan sequential search
func cetakByLokasi(T tabWisata, n int) {
	var kataKunci string
	var i, j int
	var sebagian tabWisata
	fmt.Println("Masukan kota yang ingin dicari:")
	fmt.Scan(&kataKunci)
	j = 0
	for i = 0; i < n; i++ {
		if T[i].lokasi == kataKunci {
			sebagian[j] = T[i]
			j++
		}
	}
	if j > 0 {
		printArrwisata(sebagian, j)
	} else {
		fmt.Println("Tempat wisata di kota", kataKunci, "tidak ditemukan")
	}
}

// Mencetak data berdasarkan harga yang ditemukan dengan sequential search
func cetakByHarga(T tabWisata, n int) {
	var harga int
	var i, j int
	var sebagian tabWisata
	fmt.Println("Masukan harga tempat wisata yang ingin dicari:")
	fmt.Scan(&harga)
	j = 0
	for i = 0; i < n; i++ {
		if T[i].harga == harga {
			sebagian[j] = T[i]
			j++
		}
	}
	if j > 0 {
		printArrwisata(sebagian, j)
	} else {
		fmt.Println("Harga tiket", harga, "tidak ditemukan")
	}
}

// Update data
func updateArrWisata(T *tabWisata, n, idx int) string {
	var nama, namaSebelumnya, lokasi, hari, kontak, waktuOps, pesan string
	var harga int

	namaSebelumnya = T[idx].nama
	header()
	fmt.Println("Mengubah data tempat wisata", namaSebelumnya)
	fmt.Print("Nama tempat\t: ")
	fmt.Scan(&nama)
	if findNama(*T, n, nama) != -1 {
		pesan = ("Tempat wisata " + nama + " sudah ada. ")
	} else {
		fmt.Print("Kota\t\t: ")
		fmt.Scan(&lokasi)
		fmt.Print("Harga tiket\t: ")
		fmt.Scan(&harga)
		fmt.Print("Hari buka\t: ")
		fmt.Scan(&hari)
		fmt.Print("Jam operasional\t: ")
		fmt.Scan(&waktuOps)
		fmt.Print("Kontak\t\t: ")
		fmt.Scan(&kontak)

		if konfirmasi("Ubah data?") {
			T[idx] = tempatWisata{nama, lokasi, kontak, hari, waktuOps, harga}
			pesan = "Data " + namaSebelumnya + " berhasil di diubah!"
			return pesan
		}
	}
	pesan += "Data " + namaSebelumnya + " gagal di diubah!"
	return pesan
}

// Hapus data
func hapusArrWisata(T *tabWisata, n *int, idx int) string {
	var pesan, nama string
	nama = T[idx].nama
	pesan = "Apakah anda yakin ingin menghapus data " + nama + "?"
	if konfirmasi(pesan) {
		for i := idx; i < *n; i++ {
			T[i] = T[i+1]
		}
		*n--
		return ("Data " + nama + " berhasil dihapus!")
	}
	return ("Data " + nama + " gagal dihapus!")
}

// Function konfirmasi aksi
func konfirmasi(pesan string) bool {
	var yn string
	for yn != "Y" && yn != "N" {
		fmt.Print(pesan, "(Y/N) ")
		fmt.Scan(&yn)
		if yn == "Y" || yn == "y" {
			return true
		}
	}
	return false
}

// tampilan header
func header() {
	fmt.Println("\n\n===-------------------------------------===")
	fmt.Println("            Aplikasi Pariwisata            ")
	fmt.Println("===-------------------------------------===")
}

// Halaman utama
func mainMenu(T *tabWisata, n *int) {
	var pilihan string
	var inputValid, lanjut bool = false, true

	for lanjut {
		inputValid = false
		header()
		fmt.Println("1. Lihat data")
		fmt.Println("2. Input tempat wisata")
		fmt.Println("3. Kelola tempat wisata")
		fmt.Println("0. Keluar")
		fmt.Println("=========================================")

		for !inputValid {
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case "1":
				inputValid = true
				menuPrint(*T, *n)
			case "2":
				inputValid = true
				inputArrWisata(T, n)
			case "3":
				inputValid = true
				menuKelola(T, n)
			case "0":
				lanjut = false
				inputValid = true
				fmt.Println("Terima kasih telah menggunakan aplikasi!")
			}
		}
	}

}

func menuKelola(T *tabWisata, n *int) {
	var pilihan, kataKunci, pesan, respon string
	var sebagian tabWisata
	var inputValid, lanjut bool = false, true
	var idx int

	for lanjut {
		if *n == 0 {
			fmt.Println("Data Kosong!")
			lanjut = false
		} else {
			inputValid = false
			header()
			fmt.Println("Menampilkan data tempat wisata...")
			printArrwisata(*T, *n)
			fmt.Println(respon)
			fmt.Println("\n=========================================")
			fmt.Println("1. Ubah data")
			fmt.Println("2. Hapus data")
			fmt.Println("0. Kembali")
			fmt.Println("=========================================")

			for !inputValid {
				fmt.Print("Pilihan:")
				fmt.Scan(&pilihan)

				switch pilihan {
				case "1":
					inputValid = true
					fmt.Println("Masukan nama tempat wisata yang ingin diubah:")
					fmt.Scan(&kataKunci)
					idx = findNama(*T, *n, kataKunci)
					if idx != -1 {
						sebagian[0] = T[idx]
						printArrwisata(sebagian, 1)
						pesan = "Ubah data " + sebagian[0].nama + "?"
						if konfirmasi(pesan) {
							respon = updateArrWisata(T, *n, idx)
						} else {
							respon = "Tempat wisata bernama " + kataKunci + " tidak jadi diubah"
						}
					} else {
						respon = "Tempat wisata bernama " + kataKunci + " tidak ditemukan"
					}
				case "2":
					inputValid = true
					fmt.Println("Masukan nama tempat wisata yang ingin dihapus:")
					fmt.Scan(&kataKunci)
					idx = findNama(*T, *n, kataKunci)
					if idx != -1 {
						sebagian[0] = T[idx]
						printArrwisata(sebagian, 1)
						pesan = "Hapus data " + sebagian[0].nama + "?"
						if konfirmasi(pesan) {
							respon = hapusArrWisata(T, n, idx)
						} else {
							respon = "Tempat wisata bernama " + kataKunci + " tidak jadi dihapus"
						}
					} else {
						respon = "Tempat wisata bernama " + kataKunci + " tidak ditemukan"
					}
				case "0":
					lanjut = false
					inputValid = true
				}
			}
		}
	}
}

func menuPrint(T tabWisata, n int) {
	var pilihan, kataKunci string
	var idx int
	var sebagian tabWisata
	var inputValid, lanjut bool = false, true

	for lanjut {
		if n == 0 {
			fmt.Println("Data Kosong!")
			lanjut = false
		} else {
			header()
			fmt.Println("Menampilkan data tempat wisata...")
			printArrwisata(T, n)

			for !inputValid {
				fmt.Println("\n=========================================")
				fmt.Println("1. Mencari berdasarkan nama")
				fmt.Println("2. Mencari berdasarkan lokasi")
				fmt.Println("3. Mencari berdasarkan harga")
				fmt.Println("4. Ascending nama\t5. Descending nama")
				fmt.Println("6. Ascending harga\t7. Descending harga")
				fmt.Println("0. Kembali")
				fmt.Println("=========================================")
				fmt.Print("Pilihan:")
				fmt.Scan(&pilihan)

				switch pilihan {
				case "1":
					fmt.Println("Masukan nama tempat wisata yang ingin dicari:")
					fmt.Scan(&kataKunci)
					urutDataNama(&T, n, "naik")
					idx = findNama(T, n, kataKunci)
					if idx != -1 {
						sebagian[0] = T[idx]
						printArrwisata(sebagian, 1)
					} else {
						fmt.Println("Tempat wisata bernama", kataKunci, "tidak ditemukan")
					}
				case "2":
					cetakByLokasi(T, n)
				case "3":
					cetakByHarga(T, n)
				case "4":
					fmt.Println("Data terurut ascending berdasarkan nama...")
					urutDataNama(&T, n, "naik")
					printArrwisata(T, n)
				case "5":
					fmt.Println("Data terurut descending berdasarkan nama...")
					urutDataNama(&T, n, "turun")
					printArrwisata(T, n)
				case "6":
					fmt.Println("Data terurut ascending berdasarkan harga...")
					urutDataHarga(&T, n, "naik")
					printArrwisata(T, n)
				case "7":
					fmt.Println("Data terurut descending berdasarkan harga...")
					urutDataHarga(&T, n, "turun")
					printArrwisata(T, n)
				case "0":
					lanjut = false
					inputValid = true
				}
			}
		}
	}
}
