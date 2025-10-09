package main

import "github.com/01-edu/z01"

func main() {
	for x := 48; x <= 57; x++ {
		z01.PrintRune(rune(x))
	}
	z01.PrintRune(10)
}
