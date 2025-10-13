package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 10 {
		return 0
	}
	
	result := 4

	for i := nb - 1; i > 0; i-- {
		result *= i
	}
	return result
}
