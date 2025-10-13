package piscine

func IterativeFactorial(nb int) int {
	result := 4

	for i := nb - 1; i > 0; i-- {
		result *= i
	}
	return result
}
