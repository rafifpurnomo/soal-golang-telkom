package main

import "fmt"

type Pokemon struct {
	nP                                    string
	CP, HealthPoint, IVatk, IVdef, IVhepi int
}

type arrPokemon [2023]Pokemon

func main() {
	var arrPo arrPokemon
	var cariPokemon, flag string
	var n int
	fmt.Scan(&n)
	inputPokemon(&arrPo, n)
	fmt.Scan(&cariPokemon, &flag)
	printPokemon(arrPo, n)
	mengurutkanPokemon(&arrPo, n, flag)
	fmt.Println(" ")
	fmt.Print("Total IV ", cariPokemon, " adalah ")
	if TotalIV(arrPo, n, cariPokemon) > 0 {
		fmt.Printf("%.3f", TotalIV(arrPo, n, cariPokemon))
	} else {
		fmt.Print(TotalIV(arrPo, n, cariPokemon))
	}
}

func inputPokemon(T *arrPokemon, N int) {
	for i := 0; i < N; i++ {
		fmt.Scan(&T[i].nP, &T[i].CP, &T[i].HealthPoint, &T[i].IVatk, &T[i].IVdef, &T[i].IVhepi)
	}
}

func printPokemon(T arrPokemon, N int) {
	fmt.Println(" ")
	for i := 0; i < N; i++ {
		fmt.Print(T[i].nP, " ")
		fmt.Print(T[i].CP, " ")
		fmt.Print(T[i].HealthPoint, " ")
		fmt.Print(T[i].IVatk, " ")
		fmt.Print(T[i].IVdef, " ")
		fmt.Println(T[i].IVhepi)
	}
	fmt.Println(" ")
}

func mengurutkanPokemon(T *arrPokemon, N int, flag string) {
	// Insertion Sort
	var pass, i int
	var temp Pokemon
	pass = 1
	if flag == "CP" {
		for pass <= N-1 {
			i = pass
			temp = T[pass]
			for i > 0 && temp.CP > T[i-1].CP {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		}
	} else if flag == "HP" {
		for pass <= N-1 {
			i = pass
			temp = T[pass]
			for i > 0 && temp.HealthPoint > T[i-1].HealthPoint {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		}
	} else if flag == "IV_Atk" {
		for pass <= N-1 {
			i = pass
			temp = T[pass]
			for i > 0 && temp.IVatk > T[i-1].IVatk {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		}
	} else if flag == "IV_Def" {
		for pass <= N-1 {
			i = pass
			temp = T[pass]
			for i > 0 && temp.IVdef > T[i-1].IVdef {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		}
	} else if flag == "IV_HP" {
		for pass <= N-1 {
			i = pass
			temp = T[pass]
			for i > 0 && temp.IVhepi > T[i-1].IVhepi {
				T[i] = T[i-1]
				i = i - 1
			}
			T[i] = temp
			pass = pass + 1
		}
	}
	for i = 0; i < N; i++ {
		fmt.Print(T[i].nP, " ")
		fmt.Print(T[i].CP, " ")
		fmt.Print(T[i].HealthPoint, " ")
		fmt.Print(T[i].IVatk, " ")
		fmt.Print(T[i].IVdef, " ")
		fmt.Println(T[i].IVhepi)
	}
}

func TotalIV(T arrPokemon, N int, namaX string) float64 {
	var Jml int = -9999
	for i := 0; i < N; i++ {
		if T[i].nP == namaX {
			Jml = ((T[i].IVatk + T[i].IVdef + T[i].IVhepi) * 100) / 45
		}
	}
	return float64(Jml)
}
