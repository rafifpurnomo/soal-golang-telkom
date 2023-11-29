package main

import (
	"fmt"
)

func tarif(menu int) int {
	if menu <= 3 {
		return 10000
	} else if menu > 3 && menu <= 50 {
		return 10000 + (menu-3)*2500
	} else {
		return 100000
	}

}

func hitungTarif(menu int, bersisa bool, n int, total *int) {
	if bersisa {
		*total = tarif(menu) * n
	} else {
		*total = tarif(menu)
	}
}

func main() {
	var m, jumlahMenu, banyakOrang, total int
	var sisa bool
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&jumlahMenu, &banyakOrang, &sisa)
		hitungTarif(jumlahMenu, sisa, banyakOrang, &total)
		fmt.Println(total)
	}
}
