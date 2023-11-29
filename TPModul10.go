package main

import "fmt"

const nmax int = 12345

type prodi struct {
	nama, akreditas       string
	tahun, aktif, lulusan int
}

type fakultas struct {
	nama     string
	arrprodi [nmax]prodi
	n        int
}

func buatfakultas(f *fakultas) {
	fmt.Scan(&f.nama)
}

func daftarprodi(f *fakultas) {
	var idx_prodi int
	for i := 0; i <= 9; i++ {
		fmt.Scan(&f.arrprodi[f.n].nama, &f.arrprodi[f.n].akreditas, &f.arrprodi[f.n].tahun, &f.arrprodi[f.n].aktif, &f.arrprodi[f.n].lulusan)
		idx_prodi = cekprodi(f.arrprodi[f.n].nama, *f)
		if idx_prodi >= 0 {
			f.arrprodi[idx_prodi].aktif = f.arrprodi[idx_prodi].aktif + f.arrprodi[f.n].aktif
			f.arrprodi[idx_prodi].lulusan = f.arrprodi[idx_prodi].lulusan + f.arrprodi[f.n].lulusan
			f.n = f.n - 1
		}
		f.n = f.n + 1
	}
}

func cekprodi(s string, f fakultas) int {
	var j int = 0
	for j < f.n && f.arrprodi[j].nama != s {
		j = j + 1
	}
	if j == f.n {
		return -1
	} else {
		return j
	}
}

func terbanyak(f fakultas) prodi {
	var jumlah, temp, idx int
	jumlah = f.arrprodi[0].aktif + f.arrprodi[0].lulusan
	for i := 1; i < f.n; i++ {
		temp = f.arrprodi[i].aktif + f.arrprodi[i].lulusan
		if jumlah < temp {
			jumlah = temp
			idx = i
		}
	}
	return f.arrprodi[idx]
}

func termuda(f fakultas) int {
	var idx, termuda, temp int
	termuda = f.arrprodi[0].tahun
	for i := 1; i < f.n; i++ {
		temp = f.arrprodi[i].tahun
		if temp >= termuda {
			termuda = temp
			idx = i
		}
	}
	return idx
}

func prestasi(f fakultas) string {
	var unggul, baik, cukup int
	for i := 0; i < f.n; i++ {
		if f.arrprodi[i].akreditas == "unggul" {
			unggul += 1
		} else if f.arrprodi[i].akreditas == "baik" {
			baik += 1
		} else {
			cukup += 1
		}
	}
	if unggul > baik && unggul > cukup {
		return "unggul"
	} else if baik > unggul && baik > cukup {
		return "baik"
	} else {
		return "cukup"
	}
}

func cetakprodi(f fakultas, x string) {
	x = prestasi(f)
	for i := 0; i < f.n; i++ {
		if f.arrprodi[i].akreditas == x {
			fmt.Print(f.arrprodi[i].nama, " ")
		}
	}
}
func main() {
	var f fakultas
	var x string
	var most prodi
	buatfakultas(&f)
	daftarprodi(&f)
	most = terbanyak(f)
	fmt.Printf("prodi %v memiliki mahasiswa dan lulusan terbanyak yaitu %v \n", most.nama, most.aktif+most.lulusan)
	fmt.Printf("prodi termuda adalah %v \n", f.arrprodi[termuda(f)].nama)
	x = prestasi(f)
	fmt.Printf("akreditasi Prodi terbanyak di Fakultas Informatika adalah %v \n", x)
	cetakprodi(f, x)
}

// Informatika
// IT unggul 2001 318 262
// RPL unggul 2013 488 238
// IF cukup 2019 467 305
// IT baik 2009 361 414
// SE cukup 2019 477 346
// SI cukup 2013 205 382
// DS baik 2019 224 396
// IF cukup 2016 304 268
// SD cukup 2017 296 145
// IT cukup 2007 406 267
