package main

import "fmt"

func inputBilangan(bil *int) {
	fmt.Scan(bil)
	for *bil < 0 {
		fmt.Scan(bil)
	}
}

func stop(bil int) bool {
	return bil == 0
}

func hitung(mean *float64, min, max, n *int) {
	var bilangan int
	inputBilangan(&bilangan)
	*min = bilangan
	*max = bilangan
	for !stop(bilangan) {

		if bilangan > 0 {
			*mean += float64(bilangan)
			*n++
			if bilangan < *min {
				*min = bilangan
			}
			if bilangan > *max {
				*max = bilangan
			}
		}
		inputBilangan(&bilangan)

	}

	*mean = *mean / float64(*n)
	if *n == 0 && bilangan == 0 {
		fmt.Println("- - - -")
	} else {
		fmt.Println(*mean, *max, *min, *n)
	}

}

func main() {
	var mean float64
	var min, max, n int
	hitung(&mean, &min, &max, &n)

}

// SOAL NO 2
// func average(bil, i int, hasil *float64) {
// 	var total int
// 	if bil != 0 {
// 		total = total + (bil % 10)
// 		*hasil += float64(total)
// 		average(bil/10, i+1, hasil)
// 	} else {
// 		*hasil = *hasil / float64(i)
// 	}
// }

// func main() {
// 	var hasil float64
// 	var bilangan, i int
// 	fmt.Scan(&bilangan)
// 	average(bilangan, i, &hasil)
// 	fmt.Println(hasil)
// }
