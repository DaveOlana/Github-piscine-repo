package main

import "github.com/01-edu/z01"

func main() {
	for x := 122; x >= 97; x-- {
		z01.PrintRune(rune(x))
	}
	z01.PrintRune(10)
}
