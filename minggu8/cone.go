package main

import (
	"fmt"
	"math"
)

func luasAlas(r int) float64 {
	return 3.14 * (float64(r) * float64(r))
}

func garisPelukis(r, t int) float64 {
	return math.Sqrt((float64(r) * float64(r)) + (float64(t) * float64(t)))
}

func hitungLuasSelimut(r, t int, luas *float64) {
	var s float64
	s = garisPelukis(r, t)
	*luas = 3.14 * float64(r) * s
}
func main() {
	var r1, r2, t1, t2 int
	var cone1, cone2 float64
	var conePotong float64

	fmt.Scanln(&r1, &r2, &t1, &t2)
	hitungLuasSelimut(r1, t1, &cone1)
	hitungLuasSelimut(r2, t2, &cone2)

	fmt.Printf("%3f\n", cone1)
	fmt.Printf("%3f\n", cone2)

	conePotong = luasAlas(r1) + luasAlas(r2) + (cone1 - cone2)
	fmt.Printf("%3f\n", conePotong)
}
