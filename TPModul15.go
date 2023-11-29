package main

import "fmt"

const NMAX = 100

type ASEAN struct {
	nilaiGDP [3]int
	nama     string
}

type tab_Asean [10]ASEAN

type negara [10]string

func input_negara(A *tab_Asean, list negara) {
	for i := 0; i < 10; i++ {
		A[i].nama = list[i]
	}

}
func input(A *tab_Asean) {
	var negara string
	var tahun, idxN, gdp, idxT int
	list := [10]string{"Brunei", "Cambodia", "Indonesia", "Laos", "Malaysia",
		"Myanmar", "Philippines", "Singapore", "Thailand", "Vietnam"}
	input_negara(A, list)
	fmt.Println("masukkan negara dan tahun :")
	fmt.Scan(&negara, &tahun)
	idxN = is_asean(*A, negara)
	idxT = find_tahun(tahun)
	for idxN != -1 {
		fmt.Println("masukkan gdp :")
		fmt.Scan(&gdp)
		A[idxN].nilaiGDP[idxT] = gdp
		fmt.Println("masukkan negara dan tahun :")
		fmt.Scan(&negara, &tahun)
		idxN = is_asean(*A, negara)
		idxT = find_tahun(tahun)
	}
	sort(A, idxT)
}

func print(A tab_Asean) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v %v %v %v\n", A[i].nama, A[i].nilaiGDP[0], A[i].nilaiGDP[1], A[i].nilaiGDP[2])
	}
}

func is_asean(A tab_Asean, negara string) int {
	for i := 0; i < 10; i++ {
		if A[i].nama == negara {
			return i
		}
	}
	return -1
}

func find_tahun(tahun int) int {
	if tahun == 2021 {
		return 0
	} else if tahun == 2022 {
		return 1
	} else {
		return 2
	}
}

func sort(A *tab_Asean, n int) {
	var idx_min, i, j int
	i = 1
	for i <= 10-1 {
		idx_min = i - 1
		j = i
		for j < 10 {
			if A[idx_min].nilaiGDP[n] <= A[j].nilaiGDP[n] {
				idx_min = j
			}
			j = j + 1
		}
		A[idx_min], A[i-1] = A[i-1], A[idx_min]
		i = i + 1
	}
}

func main() {
	var A tab_Asean
	input(&A)
	print(A)
}

// Vietnam 2022 615
// Malaysia 2022 163
// Indonesia 2022 884
// Philippines 2023 37
// Indonesia 2023 1893
// Philippines 2021 286
// Laos 2021 98
// Myanmar 2022 135
// Singapore 2021 357
// Vietnam 2021 638
// Thailand 2021 724
// Cambodia 2021 415
// Malaysia 2021 176
// Brunei 2023 80
// Thailand 2022 575
// Malaysia 2023 649
// Laos 2022 740
// Vietnam 2023 49
// Thailand 2023 870
// Japan 2021
