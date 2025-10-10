package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			PrintNbr(922337203685477580)
			z01.PrintRune('8')
			return
		} else if n == -2147483648 {
			PrintNbr(214748364)
			z01.PrintRune('8')
			return
		}
		n = -n
	}

	if n/10 != 0 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}
