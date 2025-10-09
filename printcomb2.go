package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a1 := '0'; a1 <= '9'; a1++ { // tens of first number
		for a2 := '0'; a2 <= '9'; a2++ { // ones of first number
			for b1 := a1; b1 <= '9'; b1++ { // tens of second number
				for b2 := '0'; b2 <= '9'; b2++ { // ones of second number
					// skip invalid cases: second number must be > first
					if b1 == a1 && b2 <= a2 {
						continue
					}
					z01.PrintRune(a1)
					z01.PrintRune(a2)
					z01.PrintRune(' ')
					z01.PrintRune(b1)
					z01.PrintRune(b2)
					if !(a1 == '9' && a2 == '8' && b1 == '9' && b2 == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
