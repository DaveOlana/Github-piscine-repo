package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 { // prevent overflow or invalid input
		return 0
	}
	if nb == 0 { // base case
		return 1
	}
	if nb > 1 {
		return nb * RecursiveFactorial(nb-1) // recursive case
	}
	return 1
}
