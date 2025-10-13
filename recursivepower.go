package piscine

func RecursivePower(base, power int) int {
	if power < 0 || power > 10 {
		return 0
	}

	if power == 0 {
		return 1
	}
	return base * RecursivePower(base, power-1)
}
