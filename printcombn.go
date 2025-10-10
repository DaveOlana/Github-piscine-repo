package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	combn := make([]int, n)
	PrintCombHelper(n, 0, 0, combn)
	z01.PrintRune(10)
}

func PrintCombHelper(n, pos, start int, combn []int) {
	if pos == n {
		PrintCombn(combn, n)
		return
	}
	for a := start; a <= 9; a++ {
		combn[pos] = a
		PrintCombHelper(n, pos+1, a+1, combn)
	}
}

func PrintCombn(combn []int, n int) {
	for a := 0; a < n; a++ {
		z01.PrintRune(rune(combn[a] + '0'))
	}
	if combn[0] != 10-n {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}
