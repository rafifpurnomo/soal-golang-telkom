package main

import (
	"fmt"
	"math"
)

func z(x, y float64) float64 {
	/* mengembalikan nilai z sesuai dengan persamaan bilangan bulat x dan y */
	return math.Sqrt(((3 * x) * (6 * y) / (4 * y)))
}

func main() {
	var a, b float64
	fmt.Scanln(&a, &b)
	fmt.Print(z(a, b), z(b, a))
}

// func jumlahBus(penumpang, kapasitas int) int {
// 	/* mengembalikan jumlah Bus yang diperlukan, apabila diketahui jumlah penumpang dan kapasitas dari bus */
// 	var bus int
// 	bus = penumpang / kapasitas
// 	if penumpang%kapasitas > 0 {
// 		bus++
// 	}
// 	return bus
// }

// func main() {
// 	var n, m int
// 	fmt.Scan(&n, &m)
// 	fmt.Println(jumlahBus(n, m), "bus")
// }

// func main() {
// 	var indeks float64
// 	var studi int
// 	var status bool
// 	fmt.Scan(&indeks, &studi, &status)
// 	if cumlaude(indeks, studi, status) {
// 		fmt.Print("Cumlaude")
// 	} else if sangatMemuaskan(indeks, studi, status) {
// 		fmt.Print("Sangat memuaskan")
// 	} else if memuaskan(indeks, studi, status) {
// 		fmt.Print("Memuaskan")
// 	}

// }

// func cumlaude(ipk float64, masaStudi int, publikasi bool) bool {
// 	/* mengembalikan true apabila memenuhi kriteria "cumlaude", dan false apabila sebaliknya */
// 	return (ipk >= 3.51 && ipk <= 4) && masaStudi <= 8 && publikasi
// }

// func sangatMemuaskan(ipk float64, masaStudi int, publikasi bool) bool {
// 	/* mengembalikan true apabila memenuhi kriteria "Sangat memuaskan", dan false apabila sebaliknya */
// 	return ipk >= 2.76 && masaStudi <= 14
// }

// func memuaskan(ipk float64, masaStudi int, publikasi bool) bool {
// 	/* mengembalikan true apabila memenuhi kriteria "Memuaskan", dan false apabila sebaliknya */
// 	return ipk >= 2 && masaStudi <= 14
// }

//  --

// func MenghitungKeliling(a, b int) int {
// 	var keliling int
// 	keliling = 2 * (a + b)
// 	return keliling
// }

// func main() {
// 	var panjang, lebar int
// 	fmt.Scan(&panjang, &lebar)
// 	fmt.Println("ini luas", MenghitungLuas(panjang, lebar))
// 	fmt.Println("ini kelilintg", MenghitungKeliling(panjang, lebar))
// }
