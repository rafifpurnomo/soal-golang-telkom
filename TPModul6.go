package main

import "fmt"

//soal no 1
// func abc_1302220003(x int) string {
// 	var (
// 		a int
// 		s string
// 	)
// 	if x != 0 {
// 		a = x % 2
// 		if a == 0 {
// 			return s + "0" + abc_1302220003(x/2)
// 		} else {
// 			return s + "1" + abc_1302220003(x/2)
// 		}
// 	}
// 	return ""
// }

// func main() {
// 	var x int
// 	fmt.Scan(&x)
// 	fmt.Print(abc_1302220003(x))
// }

// soal no 2
func beliBuku_1302220003(n, m int) {
	var sisaRakKosong int
	sisaRakKosong = n - m
	if sisaRakKosong != 1 {
		fmt.Println("beli 1 buku, sisa", sisaRakKosong-1)
		beliBuku_1302220003(n, m+1)
	} else {
		fmt.Println("beli 1 buku, rak buku penuh")
	}

}
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println("sisa rak kosong", n-m, "buku")
	beliBuku_1302220003(n, m)
}

// soal no 3
// func pola_1302220003(n int, s *string) {
// 	if n != 0 {
// 		*s = *s + "*"
// 		fmt.Println(*s)
// 		pola_1302220003(n-1, s)
// 	}
// }

// func main() {
// 	var banyakBintang int
// 	var s string
// 	fmt.Scan(&banyakBintang)
// 	pola_1302220003(banyakBintang, &s)
// }
