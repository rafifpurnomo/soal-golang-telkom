package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	display(a, b)
}

func isFaktor(num, x int) bool {
	return num%x == 0
}

func jumlahFaktor(num int, total *int) {
	*total = 0
	for i := 1; i < num; i++ {
		if isFaktor(num, i) {
			*total += i
		}
	}
}

func perfect(num int) bool {
	var hasil int
	jumlahFaktor(num, &hasil)
	return hasil == num
}

func display(x, y int) {
	fmt.Print("Barusan Perfect Number: ")
	for i := x; i <= y; i++ {
		if perfect(i) {
			fmt.Print(i, " ")
		}
	}
}
