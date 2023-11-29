package main

import "fmt"

func main() {
	var jam, menit int
	var member bool
	var total float64

	fmt.Scan(&jam, &menit, &member)
	hitungSewa(jam, menit, member, &total)
	fmt.Println(total)
}
func durasi(jam, menit int) int {
	if jam == 0 && menit != 0 {
		return 1
	} else if jam >= 1 {
		if menit > 10 {
			return jam + 1
		} else {
			return jam
		}
	} else {
		return 0
	}
}
func potongan(durasi int, tarif int) float64 {
	if durasi > 3 {
		return float64(durasi*tarif) * 10
	} else {
		return 0
	}
}
func tarif(member bool) int {
	if member {
		return 3500
	} else {
		return 5000
	}

}
func hitungSewa(jam, menit int, member bool, biaya *float64) {
	var waktuSewa int = durasi(jam, menit)
	var harga int = tarif(member)
	var diskon float64 = potongan(waktuSewa, harga)
	*biaya = float64(waktuSewa*harga) - diskon
}
