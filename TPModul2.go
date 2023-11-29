// SOAL NOMOR 1
// package main

// import "fmt"

// func main() {
// 	var x int
// 	fmt.Scan(&x)
// 	if (x%2 == 0 && x%3 == 0) || (x%3 == 0 && x%5 == 0) {
// 		fmt.Print("bilangan kelipatan 2 dan 3, atau kelipatan 3 dan 5")
// 	} else {
// 		fmt.Print("BUKAN kelipatan 2 dan 3, juga BUKAN kelipatan 3 dan 5")
// 	}
// }

// SOAL NOMOR 2
// package main

// import "fmt"

// func main() {
// 	var i, n_passed, n_failed, n int
// 	var n1, n2, n3, avg float64

// 	fmt.Print("Berapa jumlah siswa yang nilainya akan diproses?")
// 	fmt.Scan(&n)
// 	n_passed = 0
// 	n_failed = 0
// 	for i = 1; i <= n; i++ {
// 		fmt.Scan(&n1, &n2, &n3)
// 		avg = (n1 + n2 + n3) / 3
// 		if avg > 80.0 {
// 			fmt.Println("Memenuhi syarat administratif")
// 			n_passed++
// 		} else {
// 			fmt.Println("TIDAK Memenuhi syarat administratif")
// 			n_failed++
// 		}
// 	}
// 	fmt.Println("Siswa yang lolos syarat administratif", n_passed)
// 	fmt.Println("Siswa yang TIDAK lolos syarat administratif", n_failed)
// }

// SOAL NOMOR 3
package main

import (
	"fmt"
	"math"
)

func main() {
	var panjang, lebar, luas, keliling, diagonal float64
	fmt.Scan(&panjang, &lebar)
	luas = panjang * lebar
	keliling = (panjang * 2) + (lebar * 2)
	diagonal = math.Sqrt((panjang * panjang) + (lebar * lebar))
	fmt.Println("luas: ", luas)
	fmt.Println("keliling: ", keliling)
	fmt.Println("panjang diagonal: ", diagonal)
}

// SOAL NOMOR 4
// package main

// import "fmt"

// func main() {
// 	var disc, harga, total, thn int
// 	fmt.Scan(&thn, &harga)
// 	disc = (thn / 1000) * (((thn / 10) % 10 * 10) + (thn % 10))
// 	total = harga - (harga * disc / 100)
// 	fmt.Print("besar diskon: ", disc, "%\nJumlah yang dibayar: ", total)
// }
