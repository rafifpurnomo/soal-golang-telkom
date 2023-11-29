package main

import "fmt"

const NMAX int = 1000

type himpunan struct {
	info     [NMAX]string
	nElement int
}

func main() {
	var a, b, c, d himpunan
	createSet(&a)
	createSet(&b)
	intersection(a, b, &c)
	union(a, b, &d)
	printSet(c)
	fmt.Println()
	printSet(d)

}

func createSet(set *himpunan) {
	var s string
	var i int = 0
	fmt.Scan(&s)
	for !isMember(*set, s) {
		set.info[i] = s
		fmt.Scan(&s)
		i++
		set.nElement++
	}

}

func printSet(set himpunan) {
	for i := 0; i < set.nElement; i++ {
		fmt.Print(set.info[i], " ")
	}
}

func isMember(set himpunan, s string) bool {
	var i int = 0
	var member bool = false
	for !member && i < set.nElement {
		if set.info[i] == s {
			member = true
		}
		i++
	}
	return member
}

func intersection(set1, set2 himpunan, set3 *himpunan) {
	for i := 0; i <= set2.nElement; i++ {
		if isMember(set1, set2.info[i]) {
			set3.info[set3.nElement] = set2.info[i]
			set3.nElement++
		}
	}
}

func union(set1, set2 himpunan, set3 *himpunan) {
	var n int = set1.nElement
	*set3 = set1
	for i := 0; i <= set2.nElement; i++ {
		if !isMember(*set3, set2.info[i]) {
			set3.info[n] = set2.info[i]
			n++
		}
	}
	set3.nElement = n
}
