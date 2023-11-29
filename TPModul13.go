package main

import "fmt"

type student struct {
	name                            string
	math, indo, eng, sains, average float64
}

type arrStudent [2048]student

func entryStudent_1302220068(T *arrStudent, n *int) {
	var nama string
	fmt.Scan(&nama)
	for nama != "STOP" {
		T[*n].name = nama
		fmt.Scan(&T[*n].math, &T[*n].indo, &T[*n].eng, &T[*n].sains)
		T[*n].average = (T[*n].math + T[*n].indo + T[*n].eng + T[*n].sains) / 4
		*n++
		fmt.Scan(&nama)
	}
}

func max(T arrStudent, n int, flag string) int {
	var max int
	if flag == "math" {
		for i := 0; i < n; i++ {
			if T[max].math < T[i].math {
				max = i
			}
		}
	}
	if flag == "ind" {
		for i := 0; i < n; i++ {
			if T[max].indo < T[i].indo {
				max = i
			}
		}
	}
	if flag == "eng" {
		for i := 0; i < n; i++ {
			if T[max].eng < T[i].eng {
				max = i
			}
		}
	}
	if flag == "sains" {
		for i := 0; i < n; i++ {
			if T[max].sains < T[i].sains {
				max = i
			}
		}
	}
	return max
}

func ranking(T *arrStudent, n int) {

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if T[j].average < T[j+1].average {
				T[j], T[j+1] = T[j+1], T[j]
			}
		}
	}
}

func printtop3(T arrStudent, n int) {

	for i := 0; i < 3; i++ {
		fmt.Printf("Rangking %v : %v rata-rata %v \n", i+1, T[i].name, T[i].average)
	}
}

func printmax(T arrStudent, n int) {
	fmt.Printf("Nilai Matematika tertinggi diraih oleh %v \n", T[max(T, n, "math")].name)
	fmt.Printf("Nilai Bahasa Indonesia tertinggi diraih oleh %v \n", T[max(T, n, "ind")].name)
	fmt.Printf("Nilai Bahasa Inggris tertinggi diraih oleh %v \n", T[max(T, n, "eng")].name)
	fmt.Printf("Nilai pelajaran IPA/IPS tertinggi diraih oleh %v \n", T[max(T, n, "sains")].name)
}
func main() {
	var T arrStudent
	var n int
	entryStudent_1302220068(&T, &n)
	ranking(&T, n)
	printtop3(T, n)
	printmax(T, n)

}

// Ana 80 90 70 50
// Dini 10 30 90 50
// Joni 20 80 75 77
// Dika 75 71 73 60
// Desi 88 45 76 47
// Vito 67 78 67 56
// Vita 50 86 84 88
// Keysa 77 89 61 87
// Enda 87 65 77 71
