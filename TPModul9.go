// // SOAL 1
// package main

// import (
// 	"fmt"
// 	"math"
// )

// type point_1 struct { // membuat class
// 	x, y float64 // attribute class
// }

// type point_2 struct {
// 	x, y float64
// }

// func (p *point_1) setPoint_1(x, y float64) { // method setpoint 1
// 	p.x = x
// 	p.y = y
// }

// func (p *point_2) setPoint_2(xa, ya float64) { // mentthod setpoint 2
// 	p.x = xa
// 	p.y = ya
// }
// func jarak(p1 point_1, p2 point_2) float64 {
// 	return math.Sqrt(((p1.x - p2.x) * (p1.x - p2.x)) + ((p1.y - p2.y) * (p1.y - p2.y)))
// }

// func main() {
// 	var x1, x2, y1, y2 float64
// 	p1 := point_1{}
// 	p2 := point_2{}
// 	fmt.Println("silahkan untuk a dan b")
// 	fmt.Scanln(&x1, &y1, &x2, &y2)
// 	p1.setPoint_1(x1, y1)
// 	p2.setPoint_2(x2, y2)
// 	ax1, ax2, by1, by2 := p1.x, p2.x, p1.y, p2.y
// 	ab := jarak(p1, p2)
// 	fmt.Println("silahkan untuk c dan d")
// 	fmt.Scanln(&x1, &y1, &x2, &y2)
// 	p1.setPoint_1(x1, y1)
// 	p2.setPoint_2(x2, y2)
// 	cx1, cx2, dy1, dy2 := p1.x, p2.x, p1.y, p2.y
// 	cd := jarak(p1, p2)
// 	if cd < ab {
// 		fmt.Printf("%s %0.f %s %0.f %s %0.f %s %0.f %s %.1f", "Titik terdekat adalah titik C(", cx1, ",", dy1, ") dan D  C(", cx2, ",", dy2, ") adalah", cd)
// 	} else {
// 		fmt.Printf("%s %0.f %s %0.f %s %0.f %s %0.f %s %0.f", "Titik terdekat adalah titik C(", ax1, ",", by1, ") dan D  C(", ax2, ",", by2, ")  adalah", ab)
// 	}

// }

// // SOAL 2
package main

import (
	"fmt"
)

func main() {
	var cards [52]int
	var num int
	i := 0

	fmt.Print("Masukkan kartu: ")
	fmt.Scan(&num)

	for num != 0 && i < 52 {
		cards[i] = num
		i++
		fmt.Scan(&num)
	}

	fmt.Println("Kartu Amin:")
	for j := i - 1; j >= 0; j-- {
		fmt.Print(cards[j], " ")
	}
}

// // SOAL 3
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// type AsDos_T struct { // tipe bentukan
// 	nim, kode_asdos, kode_matkul, nama string
// 	jumlah_mahasiswa                   int
// }

// func (asdos *AsDos_T) inputAsdos() { // buat metode struct
// 	fmt.Print("Masukkan kode asdos : ")
// 	fmt.Scanln(&asdos.kode_asdos)
// 	fmt.Print("Masukkan nama asdos : ")
// 	fmt.Scanln(&asdos.nama)
// 	fmt.Print("Masukkan NIM asdos : ")
// 	fmt.Scanln(&asdos.nim)
// 	fmt.Print("Masukkan kode mata kuliah : ")
// 	fmt.Scanln(&asdos.kode_matkul)
// 	fmt.Print("Masukkan jumlah mahasiswa : ")
// 	fmt.Scanln(&asdos.jumlah_mahasiswa)
// }

// // parameter array mapping
// func BacaAsDos(AsDosTable *[]map[string]string, n *int) {
// 	for i := 0; i < *n; i++ {
// 		as := AsDos_T{}
// 		as.inputAsdos()
// 		// mapping atau membuat case untuk di append di line 33
// 		newRow := map[string]string{"kode_asdos": as.kode_asdos, "nim": as.nim, "nama": as.nama, "kode_Matkul": as.kode_matkul, "jumlah_Mhs": strconv.Itoa(as.jumlah_mahasiswa)}
// 		*AsDosTable = append(*AsDosTable, newRow)
// 	}
// }

// func cetakAsdos(AsDosTable []map[string]string, n int, mk string) []map[string]string { // input array 2 dimensi with return
// 	var output = []map[string]string{} // buat array 2 dimensi
// 	var mki string
// 	for i := 0; i < n; i++ {
// 		if AsDosTable[i]["kode_Matkul"] == mk { // cek jika matkul ada
// 			output = append(output, AsDosTable[i])

// 		}
// 	}
// 	if len(output) == 0 {
// 		fmt.Println("Data tidak di temukan")
// 		fmt.Println("Silahkan masukan lagi")
// 		fmt.Scanln(&mki)
// 		return cetakAsdos(AsDosTable, n, mki)
// 	}
// 	fmt.Println("Data matakuliah yang terdapat pada kode", mk)
// 	return output
// }

// func main() {
// 	var AsDosTable []map[string]string = []map[string]string{} // buat multi dimensi array
// 	var kode_matkul string
// 	var n int
// 	fmt.Println("Masukan jumlah data asdos yang ingin di masukan")
// 	fmt.Scanln(&n)
// 	BacaAsDos(&AsDosTable, &n)
// 	fmt.Println("Masukan kode matkul yang ingin di cari")
// 	fmt.Scanln(&kode_matkul)
// 	fmt.Println(cetakAsdos(AsDosTable, n, kode_matkul))
// }

// ALTERNATIF SOAL 3
// package main

// import "fmt"

// func main() {
// 	var arr TabAsDos_T
// 	var n int = 1
// 	var cariMK string
// 	BacaAsDos_1302220022(&arr, &n)
// 	fmt.Scanln(&cariMK)
// 	CetakAsDos_1302220022(arr, n, cariMK)
// }

// type AsDos_T struct {
// 	kode   string
// 	nama   string
// 	nim    int
// 	mk     string
// 	jumlah int
// }

// type TabAsDos_T [2500]AsDos_T

// func BacaAsDos_1302220022(tabel *TabAsDos_T, n *int) {
// 	for *n <= len(*tabel) || (*tabel)[*n].kode == "XXX" {
// 		fmt.Scanln(&(*tabel)[*n].kode, &(*tabel)[*n].nama, &(*tabel)[*n].nim, &(*tabel)[*n].mk, &(*tabel)[*n].jumlah)
// 		*n++
// 	}
// }

// func CetakAsDos_1302220022(tabel TabAsDos_T, n int, mk string) {
// 	for i := n - 1; i >= 1; i-- {
// 		if (tabel)[i].mk == mk {
// 			fmt.Printf("%s %s %d %s %d\n", (tabel)[i].kode, (tabel)[i].nama, (tabel)[i].nim, (tabel)[i].mk, (tabel)[i].jumlah)
// 		}
// 	}
// }
