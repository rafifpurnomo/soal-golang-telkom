// PERPUSTAKAAN
package main

import "fmt"

func inputTglPinjam(tanggal, bulan, tahun *int) {
	var tgl, bln, thn int
	fmt.Scan(&tgl, &bln, &thn)
	*tanggal = tgl
	*bulan = bln
	*tahun = thn
}

func valid(tanggal, bulan, tahun int) bool {
	var tgl int
	tgl = 0
	getJumlahHari(bulan, tahun, &tgl)
	return tahun > 0 && bulan <= 12 && bulan >= 1 && tanggal <= tgl && tanggal >= 1
}

func kabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%100 == 0 {
		return false
	} else if tahun%4 == 0 {
		return true
	} else {
		return false
	}
}

func getJumlahHari(bulan, tahun int, jmlHari *int) {
	if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		*jmlHari = 30
	} else if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		*jmlHari = 31
	} else if bulan == 2 {
		if kabisat(tahun) {
			*jmlHari = 29
		} else {
			*jmlHari = 28
		}
	}
}

func hitungTanggalKembali(tanggal1, bulan1, tahun1 int, tanggal2, bulan2, tahun2 *int) {
	var tgl int
	getJumlahHari(bulan1, tahun1, &tgl)
	if valid(tanggal1, bulan1, tahun1) {
		if tanggal1+3 > tgl {
			if bulan1 == 12 {
				*tanggal2 = tanggal1 + 3 - tgl
				*bulan2 = 1
				*tahun2 = tahun1 + 1
			} else {
				*tanggal2 = tanggal1 + 3 - tgl
				*bulan2 = bulan1 + 1
				*tahun2 = tahun1
			}
		} else {
			*tanggal2 = tanggal1 + 3
			*bulan2 = bulan1
			*tahun2 = tahun1
		}
	}

}

func main() {
	var tanggal1, bulan1, tahun1, tanggal2, bulan2, tahun2 int
	inputTglPinjam(&tanggal1, &bulan1, &tahun1)
	for valid(tanggal1, bulan1, tahun1) {
		hitungTanggalKembali(tanggal1, bulan1, tahun1, &tanggal2, &bulan2, &tahun2)
		fmt.Print(tanggal2, bulan2, tahun2)
		inputTglPinjam(&tanggal1, &bulan1, &tahun1)
	}
	fmt.Print("input tidak valid coy")
}

// BAJU KOKO
// package main

// import "fmt"

// func membeliKain(uangAwal float64, tUang *float64, tPengeluaran *float64) {
// 	*tUang = uangAwal - (uangAwal / 4)
// 	*tPengeluaran = *tPengeluaran + (uangAwal / 4)
// }
// func membeliBenangJahit(uangAwal float64, tUang *float64, tPengeluaran *float64) {
// 	*tUang = *tUang - (uangAwal / 20)
// 	*tPengeluaran = *tPengeluaran + (uangAwal / 20)
// }
// func membuatPolaBaju(uangAwal float64, tUang *float64, tPengeluaran *float64) {
// 	*tUang = *tUang - ((uangAwal / 200) * 2)
// 	*tPengeluaran = *tPengeluaran + ((uangAwal / 200) * 2)
// }
// func menjahitBaju(uangAwal float64, tUang *float64, tPengeluaran *float64) {
// 	*tUang = *tUang - ((uangAwal / 200) * 4)
// 	*tPengeluaran = *tPengeluaran + ((uangAwal / 200) * 4)
// }
// func mengemasBaju(uangAwal float64, tUang *float64, tPengeluaran *float64) {
// 	*tUang = *tUang - (((3 * uangAwal) / 200) * 2)
// 	*tPengeluaran = *tPengeluaran + (((3 * uangAwal) / 200) * 2)
// }
// func mendistribusikan(uangAwal float64, tUang *float64, tPemasukan *float64, tPengeluaran *float64) {
// 	var tUangParam, tPengeluaranParam float64
// 	membeliKain(uangAwal, &tUangParam, &tPengeluaranParam)
// 	membeliBenangJahit(uangAwal, &tUangParam, &tPengeluaranParam)
// 	membuatPolaBaju(uangAwal, &tUangParam, &tPengeluaranParam)
// 	menjahitBaju(uangAwal, &tUangParam, &tPengeluaranParam)
// 	mengemasBaju(uangAwal, &tUangParam, &tPengeluaranParam)

// 	*tPengeluaran = tPengeluaranParam + (*tPengeluaran + ((uangAwal * 3) / 200)) // distribusi
// 	*tPemasukan = uangAwal / 2
// 	*tUang = (tUangParam - +((uangAwal * 3) / 200)) + *tPemasukan
// }

// func main() {
// 	var uangAwal, tUang, tPemasukan, tPengeluaran float64
// 	fmt.Scanln(&uangAwal)
// 	mendistribusikan(uangAwal, &tUang, &tPemasukan, &tPengeluaran)
// 	fmt.Println(tPengeluaran, tPemasukan, tUang)
// }
