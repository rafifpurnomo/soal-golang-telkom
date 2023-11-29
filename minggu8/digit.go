package main

import "fmt"

func len(num int) int {
	var out int = 0
	for num > 0 {
		num = num / 10
		out++
	}
	return out
}

func pangkat(n int) int {
	// var out int = 1
	// for n > 0 {
	// 	out = out * 10
	// 	n--
	// }
	// return out

	if n > 0 {
		return 10 * pangkat(n-1)
	} else {
		return 1
	}
}

func split(num int, num1, num2 *int) {
	var tengah int = len(num) / 2
	*num1 = num / pangkat(tengah)
	*num2 = num % pangkat(tengah)
}

func main() {
	var num, num1, num2 int
	fmt.Scanln(&num)
	split(num, &num1, &num2)
	fmt.Println(num1, num2)
	fmt.Println(num1 + num2)

}
